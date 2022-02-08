package twitter

import (
	"github.com/desmos-labs/plutus/notifications"
	"github.com/desmos-labs/plutus/types"
)

var _ notifications.Client = &Client{}

// HandledPlatform implements notifications.Client
func (client *Client) App() string {
	return "twitter"
}

// SendNotification im
//// SendNotification implements notifications.Client
func (client *Client) SendNotification(data *types.Notification) error {
	//TODO implement me
	panic("implement me")
}
