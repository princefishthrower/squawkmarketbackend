package jobs

import (
	"encoding/json"
	"log"
	"os"
	"squawkmarketbackend/hub"
	"squawkmarketbackend/models"
	"squawkmarketbackend/tdameritrade"
	tdameritradeTypes "squawkmarketbackend/tdameritrade/types"
	"time"

	"github.com/gorilla/websocket"
	"github.com/philippseith/signalr"
)

func ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString string) (*websocket.Conn, *int, *tdameritradeTypes.UserPrincipals, error) {
	userPrincipals := tdameritradeTypes.UserPrincipals{}

	// marshall into a struct
	err := json.Unmarshal([]byte(userPrincipalsString), &userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// create TD Ameritrade socket connection
	conn, err := tdameritrade.CreateTDAmeritradeSocket(userPrincipals.StreamerInfo.StreamerSocketUrl)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// login to the socket connection
	requestId := 1
	err = tdameritrade.Login(requestId, *conn, userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// set express connection on the socket connection
	requestId += 1
	err = tdameritrade.SetExpressConnection(requestId, *conn, userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	return conn, &requestId, &userPrincipals, nil
}

func StartListenToSymbolsJob(server signalr.Server, est *time.Location) {
	// run in go routine
	go func() {

		userPrincipalsString := os.Getenv("TD_AMERITRADE_USER_PRINCIPALS")

		// connect to TD Ameritrade
		conn, requestId, userPrincipals, err := ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString)
		if err != nil {
			log.Println(err)
			return
		}

		// send request to listen to $SPX.X
		*requestId += 1
		err = tdameritrade.StreamSymbolQuotes("/ES", *requestId, *conn, *userPrincipals)
		if err != nil {
			log.Println(err)
			return
		}

		// logic on receiving a message
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			messageStruct := tdameritradeTypes.StreamMessage{}

			// unmarshall into a struct
			err = json.Unmarshal(message, &messageStruct)
			if err != nil {
				log.Println(err)
				return
			}

			entries := messageStruct.Data

			// loop at the entries
			for _, entry := range entries {
				if entry.Service == "QUOTE" {
					content := entry.Content
					// loop at the content
					for _, contentEntry := range content {
						// calculate the current mark - bid price is contentEntry.Two and ask price is contentEntry.Three
						markPrice := contentEntry.FortyNine

						// log the mark price to the console
						log.Println("mark price: ", markPrice)

						lastPrice := contentEntry.Three

						// now calculate percent change
						percentChange := (lastPrice - markPrice) / markPrice

						if percentChange > 0.1 {
							hub.BroadcastSquawk(server, "spx-momentum", models.Squawk{
								Feed:      "spx-momentum",
								Squawk:    "SPX momentum: 0.1 percent sub-second change",
								CreatedAt: MillisecondsToTimestampString(contentEntry.FiftyTwo),
							})
						}
					}
				}
			}
		}
	}()
	log.Println("Started Listen To Symbols Job")
}

// converts milliseconds since epoch to timestamp string
func MillisecondsToTimestampString(milliseconds int64) string {
	t := time.Unix(0, milliseconds*int64(time.Millisecond))
	return t.Format("2006-01-02 15:04:05")
}
