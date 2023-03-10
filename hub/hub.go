package hub

import (
	"fmt"
	"log"
	"squawkmarketbackend/db"
	"squawkmarketbackend/models"

	"github.com/philippseith/signalr"
)

type AppHub struct {
	signalr.Hub
}

func (h *AppHub) OnConnected(connectionID string) {
	fmt.Printf("%s connected\n", connectionID)
	h.Groups().AddToGroup("all-clients", connectionID)
}

func (h *AppHub) OnDisconnected(connectionID string) {
	fmt.Printf("%s disconnected\n", connectionID)
	h.Groups().RemoveFromGroup("all-clients", connectionID)
}

func (h *AppHub) Broadcast(message string) {
	// Broadcast to all clients
	// h.Clients().Group("all-clients").Send("freeFeedMessage", message)
}

// feed name and target event are the same
func BroadcastSquawk(server signalr.Server, feedName string, squawk models.Squawk) {
	log.Println("BROADCASTING SQUAWK TO FEED: ")
	// dumps squawk to log
	log.Println(squawk)
	server.HubClients().Group(feedName).Send(feedName, squawk)
}

func (h *AppHub) AddToGroup(group string, connectionID string) {
	log.Println("PERSON JOINING! " + connectionID)
	log.Println("GROUP: " + group)
	h.Groups().AddToGroup(group, connectionID)

	// send them the latest squawk by the feed they just subscribed to
	squawk, err := db.GetLatestSquawkByFeed(group)
	if err != nil {
		fmt.Printf("Error getting latest squawk: %s", err)
		return
	}

	// send the squawk only to the client that just connected
	h.Clients().Client(connectionID).Send(group, squawk)
}

func (h *AppHub) RemoveFromGroup(group string, connectionID string) {
	h.Groups().RemoveFromGroup(group, connectionID)
}
