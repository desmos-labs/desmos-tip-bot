package streamlabs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/desmos-labs/desmostipbot/database"
	"github.com/desmos-labs/desmostipbot/types"
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

// tokenRequestResponse represents the content of the token request response body
type tokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Error        string `json:"error"`
	Message      string `json:"message"`
}

// RefreshAuthenticationToken retrieves the authentication code for the user with the given
// Desmos address using the provided oAuth code. Once done, the access token and refresh tokens
// are both stored inside the database.
func (client *Client) RefreshAuthenticationToken(desmosAddress string, oAuthCode string) error {
	// Build the params values
	params := url.Values{}
	params.Add("client_id", client.clientID)
	params.Add("client_secret", client.clientSecret)
	params.Add("redirect_uri", client.redirectURI)
	params.Add("grant_type", "authorization_code")
	params.Add("code", oAuthCode)

	// Get the authentication token
	res, err := http.PostForm("https://streamlabs.com/api/v1.0/token", params)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	var resBody tokenRequestResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return err
	}

	if resBody.Error != "" {
		return fmt.Errorf("%s-%s", resBody.Error, resBody.Message)
	}

	// Store the token inside the database
	oAuthToken := types.NewOAuthToken(desmosAddress, types.ServiceStreamlabs, resBody.AccessToken, resBody.RefreshToken)
	err = client.db.SaveOAuthToken(oAuthToken)
	if err != nil {
		return err
	}

	return err
}
