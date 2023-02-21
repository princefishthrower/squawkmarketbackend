package hub

import (
	"fmt"
	"squawkmarketbackend/db"
	scraperTypes "squawkmarketbackend/scraper/types"

	"github.com/philippseith/signalr"
)

type AppHub struct {
	signalr.Hub
}

func (h *AppHub) OnConnected(connectionID string) {
	fmt.Printf("%s connected\n", connectionID)
	h.Groups().AddToGroup("group", connectionID)

	// also send them the latest headline
	headline, err := db.GetLatestHeadline()
	if err != nil {
		fmt.Printf("Error getting latest headline: %s", err)
		return
	}

	// send the headline only to the client that just connected
	h.Clients().Client(connectionID).Send("freeFeedMessage", headline)
}

func (h *AppHub) OnDisconnected(connectionID string) {
	fmt.Printf("%s disconnected\n", connectionID)
	h.Groups().RemoveFromGroup("group", connectionID)
}

func (h *AppHub) Broadcast(message string) {
	// Broadcast to all clients
	h.Clients().Group("group").Send("freeFeedMessage", message)
}

func BroadcastHeadline(headline scraperTypes.Headline, server signalr.Server) {
	server.HubClients().All().Send("freeFeedMessage", headline)
}
