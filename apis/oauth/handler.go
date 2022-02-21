package oauth

import (
	"github.com/desmos-labs/plutus/types"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/oauth/handler"
)

// Handler represents the class that allows to handle various requests
type Handler struct {
	cfg     *types.DonationsConfig
	db      *database.Database
	cdc     codec.Codec
	amino   *codec.LegacyAmino
	handler *handler.OAuthHandler
}

// NewHandler returns a new Handler instance
func NewHandler(cfg *types.DonationsConfig, handler *handler.OAuthHandler, cdc codec.Codec, amino *codec.LegacyAmino, db *database.Database) *Handler {
	return &Handler{
		cfg:     cfg,
		db:      db,
		cdc:     cdc,
		amino:   amino,
		handler: handler,
	}
}

// HandleAuthenticationTokenRequest allows to handle a request for an authentication token
func (h *Handler) HandleAuthenticationTokenRequest(request TokenRequest) error {
	// Verify the request
	err := request.Verify(request.OAuthCode, h.cdc, h.amino)
	if err != nil {
		return err
	}

	// Get the user account
	user := types.NewUser(request.DesmosAddress)

	// Get the service account
	serviceAccount, err := h.handler.GetServiceAccount(request.Platform, request.OAuthCode)
	if err != nil {
		return err
	}

	// Get all the application accounts
	var applications []*types.ApplicationAccount
	for _, application := range h.cfg.SupportedApps {
		// Get the application username
		username, err := h.handler.GetApplicationUsername(request.Platform, application, serviceAccount)
		if err != nil {
			return err
		}

		if username == "" {
			// Skip if the username was not found
			continue
		}

		// Build the application object
		applications = append(applications, types.NewApplicationAccount(application, username))
	}

	// Store the token in the database
	return h.db.SaveUserData(user, serviceAccount, applications)
}
