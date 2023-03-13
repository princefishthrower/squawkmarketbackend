package tdameritrade

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"time"

	tdameritradeTypes "squawkmarketbackend/tdameritrade/types"

	"github.com/gorilla/websocket"
)

func RefreshStreamToken() {

}

func CreateTDAmeritradeSocket(streamerSocketUrl string) (*websocket.Conn, error) {
	addr := streamerSocketUrl
	url := url.URL{Scheme: "wss", Host: addr, Path: "/ws"}
	log.Printf("connecting to %s", url.String())

	websocketConn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
		return nil, err
	}

	return websocketConn, nil
}

func ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString string) (*websocket.Conn, *int, *tdameritradeTypes.UserPrincipals, error) {
	userPrincipals := tdameritradeTypes.UserPrincipals{}

	// marshall into a struct
	err := json.Unmarshal([]byte(userPrincipalsString), &userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// create TD Ameritrade socket connection
	conn, err := CreateTDAmeritradeSocket(userPrincipals.StreamerInfo.StreamerSocketUrl)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// login to the socket connection
	requestId := 1
	err = Login(requestId, *conn, userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// set express connection on the socket connection
	requestId += 1
	err = SetExpressConnection(requestId, *conn, userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	return conn, &requestId, &userPrincipals, nil
}

func Login(requestID int, conn websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {

	tokenTimeStampAsDateObj, err := time.Parse("2006-01-02T15:04:05-0700", userPrincipals.StreamerInfo.TokenTimestamp)
	if err != nil {
		log.Println(err)
		return err
	}
	tokenTimeStampAsMs := tokenTimeStampAsDateObj.UnixNano() / int64(time.Millisecond)

	credentials := tdameritradeTypes.Credentials{
		UserId:      userPrincipals.Accounts[0].AccountId,
		Token:       userPrincipals.StreamerInfo.Token,
		Company:     userPrincipals.Accounts[0].Company,
		Segment:     userPrincipals.Accounts[0].Segment,
		CdDomain:    userPrincipals.Accounts[0].AccountCdDomainId,
		UserGroup:   userPrincipals.StreamerInfo.UserGroup,
		AccessLevel: userPrincipals.StreamerInfo.AccessLevel,
		Authorized:  "Y",
		Timestamp:   tokenTimeStampAsMs,
		AppID:       userPrincipals.StreamerInfo.AppID,
		ACL:         userPrincipals.StreamerInfo.ACL,
	}

	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			tdameritradeTypes.WebSocketRequest{
				Service:   "ADMIN",
				Command:   "LOGIN",
				RequestID: requestID,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"credential": jsonToQueryString(credentials),
					"token":      userPrincipals.StreamerInfo.Token,
					"version":    "1.0",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func SetExpressConnection(requestID int, conn websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			tdameritradeTypes.WebSocketRequest{
				Service:   "ADMIN",
				Command:   "QOS",
				RequestID: requestID,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"qoslevel": "0",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil

}

func StreamSymbolQuotes(symbol string, requestID int, conn websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			tdameritradeTypes.WebSocketRequest{
				Service:   "QUOTE",
				Command:   "SUBS",
				RequestID: 1,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"keys":   symbol,
					"fields": "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func jsonToQueryString(json tdameritradeTypes.Credentials) string {
	values := url.Values{}

	v := reflect.ValueOf(json)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("json")
		value := fmt.Sprintf("%v", v.Field(i))
		values.Add(tag, value)
	}

	return values.Encode()
}
