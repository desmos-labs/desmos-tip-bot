package donations

import (
	"fmt"
	"net/http"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/desmos"
	"github.com/desmos-labs/plutus/notifications/handler"
	"github.com/desmos-labs/plutus/types"
)

// Handler allows to handle the requests related to the donations
type Handler struct {
	db     *database.Database
	desmos *desmos.Client

	notificationsHandler *handler.NotificationsHandler
}

// NewHandler returns a new Handler instance
func NewHandler(
	desmosClient *desmos.Client,
	notificationsHandler *handler.NotificationsHandler,
	db *database.Database,
) *Handler {
	return &Handler{
		desmos:               desmosClient,
		db:                   db,
		notificationsHandler: notificationsHandler,
	}
}

// HandleDonationRequest handles the given request trying to perform the donation
func (h *Handler) HandleDonationRequest(request DonationRequest) error {
	// Validity checks
	if len(request.TipperUsername) > 25 {
		return apiutils.WrapErr(http.StatusBadRequest, fmt.Sprintf(
			"Tipper name is too long. Max: 25, got %d", len(request.TipperUsername)))
	}

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

	// Get the donor profile
	profile, err := h.desmos.GetDesmosProfile(donationTx.SenderAddress)
	if err != nil {
		return err
	}

	// Send the notification
	return h.notificationsHandler.SendNotification(types.NewNotification(
		appAccount,
		donationTx,
		profile,
		request.TipperUsername,
		request.DonationMessage,
	))
}
