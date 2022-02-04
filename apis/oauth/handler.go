package oauth

import (
	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/oauth/handler"
)

// Handler represents the class that allows to handle various requests
type Handler struct {
	db      *database.Database
	handler *handler.OAuthHandler
}

// NewHandler returns a new Handler instance
func NewHandler(handler *handler.OAuthHandler, db *database.Database) *Handler {
	return &Handler{
		db:      db,
		handler: handler,
	}
}

// HandleAuthenticationTokenRequest allows to handle a request for an authentication token
func (h *Handler) HandleAuthenticationTokenRequest(service, desmosAddress, oAuthCode string) error {
	// Get the token
	token, err := h.handler.GetAuthenticationToken(service, desmosAddress, oAuthCode)
	if err != nil {
		return err
	}

	// Store the token in the database
	return h.db.SaveOAuthToken(token)
}
