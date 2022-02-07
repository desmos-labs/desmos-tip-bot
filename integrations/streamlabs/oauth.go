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

type badRequestResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// getBadRequestError returns the error to be returned if the given response has returned status 400
func (client *Client) getBadRequestError(res *http.Response) error {
	if res.StatusCode != http.StatusBadRequest {
		return nil
	}

	var body badRequestResponse
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return err
	}

	return fmt.Errorf("%s-%s", body.Error, body.Message)
}

// applicationData contains the data of a single application connected to Streamlabs
type applicationData struct {
	Name string `json:"name"`
}

// GetApplicationUsername implements Client
func (client *Client) GetApplicationUsername(application string, token *types.ServiceAccount) (string, error) {
	req, err := http.NewRequest("GET", "https://streamlabs.com/api/v1.0/user", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("access_token", token.AccessToken)
	req.URL.RawQuery = q.Encode()

	res, err := client.http.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	err = client.getBadRequestError(res)
	if err != nil {
		return "", err
	}

	var resBody map[string]applicationData
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return "", err
	}

	data, ok := resBody[application]
	if !ok {
		return "", nil
	}

	return data.Name, nil
}

// tokenRequestResponse represents the content of the token request response body
type tokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GetAuthenticationToken implements Client
func (client *Client) GetServiceAccount(oAuthCode string) (*types.ServiceAccount, error) {
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

	err = client.getBadRequestError(res)
	if err != nil {
		return nil, err
	}

	// Read the response body
	var resBody tokenRequestResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return nil, err
	}

	// Get the Twitch username

	// Store the token inside the database
	return types.NewServiceAccount(
		client.Service(),
		resBody.AccessToken,
		resBody.RefreshToken,
	), nil
}

// RefreshToken implements Client
func (client *Client) RefreshToken(token *types.ServiceAccount) (*types.ServiceAccount, error) {
	//TODO implement me
	panic("implement me")
}
