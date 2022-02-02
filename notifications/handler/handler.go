package handler

import (
	"strings"

	"github.com/desmos-labs/desmostipbot/notifications"
)

var _ notifications.Client = &NotificationsHandler{}

// NotificationsHandler allows handling notifications by calling the proper client
type NotificationsHandler struct {
	clients []notifications.Client
}

// NewNotificationsHandler returns a new NotificationsHandler
func NewNotificationsHandler() *NotificationsHandler {
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
func (n *NotificationsHandler) SendNotification(data *notifications.Data) error {
	for _, client := range n.clients {
		if strings.EqualFold(client.HandledPlatform(), data.Platform) {
			err := client.SendNotification(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
