package streamlabs

import (
	"github.com/desmos-labs/desmostipbot/integrations/streamlabs"
)

// Handler represents the class that allows to handle various requests
type Handler struct {
	streamlabsClient *streamlabs.Client
}

// NewHandler returns a new Handler instance
func NewHandler(client *streamlabs.Client) *Handler {
	return &Handler{
		streamlabsClient: client,
	}
}

// HandleAuthenticationTokenRequest allows to handle a request for an authentication token
func (h *Handler) HandleAuthenticationTokenRequest(desmosAddress string, oAuthCode string) error {
	return h.streamlabsClient.RefreshAuthenticationToken(desmosAddress, oAuthCode)
}
