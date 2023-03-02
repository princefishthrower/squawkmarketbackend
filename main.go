package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"squawkmarketbackend/hub"
	"squawkmarketbackend/jobs"
	"squawkmarketbackend/stripe"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/philippseith/signalr"
)

func main() {
	// read env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// run up migrations
	RunUpMigrations()

	// run server
	RunServer()
}

func RunServer() {
	hub := &hub.AppHub{}

	// build a signalr.Server using your hub
	// and any server options you may need
	server, _ := signalr.NewServer(context.Background(),
		signalr.SimpleHubFactory(hub),
		signalr.KeepAliveInterval(2*time.Second),
		signalr.Logger(kitlog.NewLogfmtLogger(os.Stderr), true),
		signalr.InsecureSkipVerify(false),
		signalr.AllowOriginPatterns([]string{os.Getenv("EXTERNAL_URL")}),
	)

	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)
		return
	}

	// start scraping job using the server
	jobs.StartHeadlineScrapeJob(server)

	// economic prints
	jobs.EconomicPrintScrapeJob(server)

	// finviz scrape job
	jobs.StartFinvizScrapeJob(server)

	// google custom search job
	// TODO: figure out google news bug here
	// TODO: figure out costs here
	// jobs.StartGoogleCustomSearchJob(server)

	// premarket job
	jobs.StartPremarketJob(server, est)

	// post market job
	jobs.StartPostmarketJob(server, est)

	// heartbeat job for uptime
	jobs.StartHeartBeatJob()

	// delete squawks job
	jobs.DeleteSquawksJob(server, est)

	// market open job
	jobs.StartMarketOpenJob(server, est)

	// market closed job
	jobs.StartMarketClosedJob(server, est)

	// create a new http.ServerMux to handle your app's http requests
	router := http.NewServeMux()

	// ask the signalr server to map it's server
	// api routes to your custom baseurl
	server.MapHTTP(signalr.WithHTTPServeMux(router), "/feeds")

	// add stripe route
	router.HandleFunc("/handle-stripe-webhook", stripe.HandleStripeWebhook)

	log.Printf("Server starting, %s", os.Getenv("SERVER_URL"))
	err = http.ListenAndServe(os.Getenv("SERVER_URL"), LogRequests(router))
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func LogRequests(h http.Handler) http.Handler {
	// type our middleware as an http.HandlerFunc so that it is seen as an http.Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// log all headers in the request
		// for name, headers := range r.Header {
		// 	for _, h := range headers {
		// 		log.Printf("I SEE HEADER: %v: %v", name, h)
		// 	}
		// }

		// log origin
		log.Printf("I SEE ORIGIN: %v", r.Header.Get("Origin"))

		// if the origin is the staging site, allow CORS for it
		corsSite := os.Getenv("EXTERNAL_URL")
		if r.Header.Get("Origin") == "https://staging.squawk-market.com" {
			corsSite = "https://staging.squawk-market.com"
		}

		// sample CORS handling
		w.Header().Set("Access-Control-Allow-Origin", corsSite)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization,X-Requested-With,X-SignalR-User-Agent")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// serve the inner request
		h.ServeHTTP(w, r)
	})
}

func RunUpMigrations() {
	// up migrations
	// Open a database connection
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create the table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS squawks (
		id INTEGER PRIMARY KEY,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		link TEXT,
		symbols TEXT,
		feed TEXT NOT NULL,
		squawk TEXT NOT NULL,
		mp3data BLOB NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Up migrations run successfully.")
}
