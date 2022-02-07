package oauth

import "github.com/desmos-labs/plutus/types"

// Client represents a generic oAuth client
type Client interface {
	// Service returns the service that this client refers to
	Service() string

	// GetAuthenticationToken allows to get the first authorization token associated with
	// the provided OAuth code
	GetAuthenticationToken(oAuthCode string) (*types.ServiceAccount, error)

	// GetApplicationUsername allows to get the username on the specified application for
	// the user with the provided token
	GetApplicationUsername(application string, token *types.ServiceAccount) (string, error)

	// RefreshToken allows to refresh the given token by returning a new one
	RefreshToken(token *types.ServiceAccount) (*types.ServiceAccount, error)
}
