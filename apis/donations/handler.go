package donations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmostipbot/client"
	"github.com/desmos-labs/desmostipbot/notifications"
)

// Handler allows to handle the requests related to the donations
type Handler struct {
	desmosClient        *client.DesmosClient
	notificationsClient notifications.Client
}

// NewHandler returns a new Handler instance
func NewHandler(client *client.DesmosClient, notificationsClient notifications.Client) *Handler {
	return &Handler{
		desmosClient:        client,
		notificationsClient: notificationsClient,
	}
}

// HandleDonationRequest handles the given request trying to perform the donation
func (h *Handler) HandleDonationRequest(request DonationRequest) (txHash string, err error) {
	// Get the tipper address
	tipperAddress, err := h.desmosClient.ParseAddress(request.TipperAddress)
	if err != nil {
		return "", err
	}

	// Get the recipient address
	recipientAddress, err := h.desmosClient.SearchDesmosAddress(request.Platform, request.Username)
	if err != nil {
		return "", err
	}

	amount, err := sdk.ParseCoinsNormalized(request.Amount)
	if err != nil {
		return "", err
	}

	// Try sending the tip
	txHash, err = h.desmosClient.SendTip(tipperAddress, amount, recipientAddress)
	if err != nil {
		return "", err
	}

	// Send the notification to the recipient
	err = h.notificationsClient.SendNotification(notifications.NewData(
		request.Platform,
		request.Username,
		recipientAddress.String(),
		request.DonationMessage,
		txHash,
	))

	return txHash, err
}
