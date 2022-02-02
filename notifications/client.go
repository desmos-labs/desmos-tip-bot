package notifications

// Client represents a generic client that allows to send out notifications
type Client interface {
	// HandledPlatform returns the identifier of the platform that this client can handle notifications for
	HandledPlatform() string

	// SendNotification sends the notification for the given data
	SendNotification(data *Data) error
}
