package notifications

import "github.com/desmos-labs/plutus/types"

// Client represents a generic client that allows to send out notifications
type Client interface {
	// App returns the name of the application that this client can handle notifications for (Twitch, Twitter, etc).
	App() string

	// SendNotification sends the notification for the given data.
	// Prerequisites:
	// - The given donation tx has been successful
	// - The donation recipient has a Desmos profile connected to the given application account
	SendNotification(data *types.Notification) error
}
