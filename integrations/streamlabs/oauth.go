package streamlabs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/desmos-labs/plutus/oauth"
	"github.com/desmos-labs/plutus/types"
)

var _ oauth.Client = &Client{}

func (client *Client) Service() string {
	return "Streamlabs"
}

// tokenRequestResponse represents the content of the token request response body
type tokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Error        string `json:"error"`
	Message      string `json:"message"`
}

func (client *Client) GetAuthenticationToken(desmosAddress string, oAuthCode string) (*types.OAuthToken, error) {
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
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body
	var resBody tokenRequestResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return nil, err
	}

	if resBody.Error != "" {
		return nil, fmt.Errorf("%s-%s", resBody.Error, resBody.Message)
	}

	// Store the token inside the database
	return types.NewOAuthToken(
		desmosAddress,
		client.Service(),
		resBody.AccessToken,
		resBody.RefreshToken,
	), nil
}

func (client *Client) RefreshToken(token *types.OAuthToken) (*types.OAuthToken, error) {
	//TODO implement me
	panic("implement me")
}
