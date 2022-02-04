package streamlabs

import (
	"github.com/desmos-labs/plutus/notifications"
	"github.com/desmos-labs/plutus/types"
)

var _ notifications.Client = &Client{}

// HandledPlatform implements notifications.Client
func (client *Client) HandledPlatform() string {
	return "twitch"
}

// SendNotification implements notifications.Client
func (client *Client) SendNotification(data *types.Notification) error {
	// Get the authorization token
	token, err := client.db.GetOAuthToken(data.DesmosAddress)
	if err != nil {
		return err
	}

	// If the token does not exist, simply return
	if token == nil {
		return nil
	}

	// TODO: Refresh the token
	// TODO: Send the alert notification

	return nil
}
