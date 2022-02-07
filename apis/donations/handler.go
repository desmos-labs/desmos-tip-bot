package donations

import (
	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/desmos"
	"github.com/desmos-labs/plutus/notifications/handler"
	"github.com/desmos-labs/plutus/types"
)

// Handler allows to handle the requests related to the donations
type Handler struct {
	desmos *desmos.Client
	db     *database.Database

	notificationsHandler *handler.NotificationsHandler
}

// NewHandler returns a new Handler instance
func NewHandler(client *desmos.Client, notificationsHandler *handler.NotificationsHandler, db *database.Database) *Handler {
	return &Handler{
		desmos:               client,
		db:                   db,
		notificationsHandler: notificationsHandler,
	}
}

// HandleDonationRequest handles the given request trying to perform the donation
func (h *Handler) HandleDonationRequest(request DonationRequest) error {
	// Get the app account
	appAccount, err := h.db.GetAppAccount(request.RecipientApplication, request.RecipientUsername)
	if err != nil {
		return err
	}

	if appAccount == nil {
		// No app account to send the notification to
		return nil
	}

	// Get the donation tx
	donationTx, err := h.desmos.GetDonationDetails(request.TxHash)
	if err != nil {
		return err
	}

	// Send the notification
	notification := types.NewNotification(appAccount, donationTx, request.TipperUsername, request.DonationMessage)
	return h.notificationsHandler.SendNotification(notification)
}
