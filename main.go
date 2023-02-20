package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"squawkmarketbackend/hub"
	"squawkmarketbackend/jobs"
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

	hub := &hub.AppHub{}

	// build a signalr.Server using your hub
	// and any server options you may need
	server, _ := signalr.NewServer(context.TODO(),
		signalr.SimpleHubFactory(hub),
		signalr.KeepAliveInterval(2*time.Second),
		signalr.Logger(kitlog.NewLogfmtLogger(os.Stderr), true),
	)

	// start headline scrape job using the server
	jobs.StartHeadlineScrapeJob(server)

	// create a new http.ServerMux to handle your app's http requests
	router := http.NewServeMux()

	// ask the signalr server to map it's server
	// api routes to your custom baseurl
	server.MapHTTP(signalr.WithHTTPServeMux(router), "/feed")

	if err := http.ListenAndServe(os.Getenv("SERVER_URL"), router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	// http.HandleFunc("/stream", websocket.HandleWebsocket)
	// http.ListenAndServe(":8080", nil)

	// migrations - TODO: get CLI way to work

	// // Open a database connection
	// db, err := sql.Open("sqlite3", "squawkmarketbackend.db")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// // Create the headlines table
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS headlines (
	// 	id INTEGER PRIMARY KEY,
	// 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	// 	headline TEXT NOT NULL,
	// 	mp3data BLOB NOT NULL
	// )`)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Database and table initialized successfully.")
}
