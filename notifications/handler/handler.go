package handler

import (
	"strings"

	"github.com/desmos-labs/plutus/types"

	"github.com/desmos-labs/plutus/notifications"
)

// NotificationsHandler allows handling notifications by calling the proper client
type NotificationsHandler struct {
	clients []notifications.Client
}

// NewHandler returns a new NotificationsHandler
func NewHandler() *NotificationsHandler {
	return &NotificationsHandler{}
}

// RegisterClient registers the given client inside the list of clients to be called
func (n *NotificationsHandler) RegisterClient(client notifications.Client) *NotificationsHandler {
	n.clients = append(n.clients, client)
	return n
}

// HandledPlatform implements notifications.Client
func (n *NotificationsHandler) HandledPlatform() string {
	return "none"
}

// SendNotification implements notifications.Client
func (n *NotificationsHandler) SendNotification(data *types.Notification) error {
	for _, client := range n.clients {
		if strings.EqualFold(client.HandledPlatform(), data.RecipientAppAccount.Application) {
			err := client.SendNotification(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
