package notifications

// Data contains all the data that might be needed in order to send a notification
type Data struct {
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

func NewData(platform, username, desmosAddress, message, txHash string) *Data {
	return &Data{
		Platform:      platform,
		Username:      username,
		DesmosAddress: desmosAddress,
		Message:       message,
		TxHash:        txHash,
	}
}
