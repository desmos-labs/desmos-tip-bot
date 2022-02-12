package config

import "github.com/desmos-labs/plutus/desmos"

// Handler allows to handle config related requests
type Handler struct {
	desmos *desmos.Client
}

// NewHandler returns a new Handler instance
func NewHandler(desmos *desmos.Client) *Handler {
	return &Handler{
		desmos: desmos,
	}
}

// Response contains the data that is returned after a config request
type Response struct {
	Wallet string `json:"wallet"`
}

// HandleRequestConfig handles the request of a config, returning a new Response instance
func (h *Handler) HandleRequestConfig() *Response {
	return &Response{
		Wallet: h.desmos.GetAddress(),
	}
}
