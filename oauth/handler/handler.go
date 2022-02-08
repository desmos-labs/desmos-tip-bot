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

// GetServiceAccount implements oauth.Client
func (h *OAuthHandler) GetServiceAccount(service, oAuthCode string) (*types.ServiceAccount, error) {
	for _, c := range h.clients {
		if strings.EqualFold(c.Service(), service) {
			return c.GetServiceAccount(oAuthCode)
		}
	}

	return nil, fmt.Errorf("invalid service type: %s", service)
}

// GetApplicationUsername implements oauth.Client
func (h *OAuthHandler) GetApplicationUsername(service, application string, token *types.ServiceAccount) (string, error) {
	for _, c := range h.clients {
		if strings.EqualFold(c.Service(), service) {
			return c.GetApplicationUsername(application, token)
		}
	}

	return "", nil
}

// RefreshToken implements oauth.Client
func (h *OAuthHandler) RefreshToken(token *types.ServiceAccount) (*types.ServiceAccount, error) {
	for _, c := range h.clients {
		if strings.EqualFold(c.Service(), token.Service) {
			return c.RefreshAccount(token)
		}
	}

	return nil, fmt.Errorf("invalid service type: %s", token.Service)
}
