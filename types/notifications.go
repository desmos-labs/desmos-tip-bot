package types

// Notification contains all the data that might be needed in order to send a notification
type Notification struct {
	// Platform used to send the donation
	Platform string

	// Username of the recipient of the donation
	Username string

	// Desmos address of the recipient of the donation
	DesmosAddress string

	// Donation message
	Message string

	// Hash of the transaction sent for the donation
	TxHash string
}

// NewNotification returns a new Notification instance
func NewNotification(platform, username, desmosAddress, message, txHash string) *Notification {
	return &Notification{
		Platform:      platform,
		Username:      username,
		DesmosAddress: desmosAddress,
		Message:       message,
		TxHash:        txHash,
	}
}
