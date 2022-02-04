package handler

import (
	"fmt"
	"strings"

	"github.com/desmos-labs/plutus/oauth"
	"github.com/desmos-labs/plutus/types"
)

// OAuthHandler represents a handler for all oAuth-related requests.
// Underlying it relies on a series of OAuth clients that are able to fetch and refresh tokens properly.
type OAuthHandler struct {
	clients []oauth.Client
}

// NewHandler returns a new OAuthHandler instance
func NewHandler() *OAuthHandler {
	return &OAuthHandler{}
}

// RegisterClient registers the given client inside the list of clients to be called
func (h *OAuthHandler) RegisterClient(client oauth.Client) *OAuthHandler {
	h.clients = append(h.clients, client)
	return h
}

// Service implements oauth.Client
func (h *OAuthHandler) Service() string {
	return "none"
}

// GetAuthenticationToken implements oauth.Client
func (h *OAuthHandler) GetAuthenticationToken(service, desmosAddress, oAuthCode string) (*types.OAuthToken, error) {
	for _, c := range h.clients {
		if strings.EqualFold(c.Service(), service) {
			return c.GetAuthenticationToken(desmosAddress, oAuthCode)
		}
	}

	return nil, fmt.Errorf("invalid service type: %s", service)
}

// RefreshToken implements oauth.Client
func (h *OAuthHandler) RefreshToken(token *types.OAuthToken) (*types.OAuthToken, error) {
	for _, c := range h.clients {
		if strings.EqualFold(c.Service(), token.Service) {
			return c.RefreshToken(token)
		}
	}

	return nil, fmt.Errorf("invalid service type: %s", token.Service)
}
