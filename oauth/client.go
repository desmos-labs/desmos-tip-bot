package oauth

import "github.com/desmos-labs/plutus/types"

// Client represents a generic oAuth client
type Client interface {
	// Service returns the service that this client refers to
	Service() string

	// GetServiceAccount allows to get the first service account associated with
	// the provided OAuth code
	GetServiceAccount(oAuthCode string) (*types.ServiceAccount, error)

	// GetApplicationUsername allows to get the username on the specified application for
	// the user with the provided token
	GetApplicationUsername(application string, token *types.ServiceAccount) (string, error)

	// RefreshAccount allows to refresh the given service account by returning a new one
	RefreshAccount(token *types.ServiceAccount) (*types.ServiceAccount, error)
}
