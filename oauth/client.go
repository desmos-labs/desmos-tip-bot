package oauth

import "github.com/desmos-labs/plutus/types"

// Client represents a generic oAuth client
type Client interface {
	// Service returns the service that this client refers to
	Service() string

	// GetAuthenticationToken allows to get the first authorization token for the user with the given
	// Desmos address using the provided oAuth code
	GetAuthenticationToken(desmosAddress string, oAuthCode string) (*types.OAuthToken, error)

	// RefreshToken allows to refresh the given token by returning a new one
	RefreshToken(token *types.OAuthToken) (*types.OAuthToken, error)
}
