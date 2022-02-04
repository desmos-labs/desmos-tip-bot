package streamlabs

import (
	"net/http"

	"github.com/desmos-labs/plutus/database"
	"github.com/desmos-labs/plutus/types"
)

// Client represents the client used to interact with Streamlabs APIs
type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string

	http *http.Client
	db   *database.Database
}

// NewClient returns a new Client instance
func NewClient(cfg *types.StreamlabsConfig, db *database.Database) *Client {
	return &Client{
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		redirectURI:  cfg.RedirectURI,
		http:         &http.Client{},
		db:           db,
	}
}
