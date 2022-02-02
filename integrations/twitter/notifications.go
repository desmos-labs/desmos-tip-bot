package twitter

import (
	"github.com/desmos-labs/desmostipbot/notifications"
)

var _ notifications.Client = &Client{}

// HandledPlatform implements notifications.Client
func (client *Client) HandledPlatform() string {
	return "twitter"
}

// SendNotification im
//// SendNotification implements notifications.Client
func (client *Client) SendNotification(data *notifications.Data) error {
	//TODO implement me
	panic("implement me")
}
