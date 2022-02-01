package streamlabs

import (
	"encoding/json"
	"github.com/desmos-labs/desmostipbot/database"
	"github.com/desmos-labs/desmostipbot/types"
	"net/http"
	"net/url"
)

type Handler struct {
	cfg *types.StreamlabsConfig

	client *http.Client
	db     *database.Database
}

func NewHandler(cfg *types.StreamlabsConfig, db *database.Database) *Handler {
	return &Handler{
		cfg:    cfg,
		client: &http.Client{},
		db:     db,
	}
}

// --------------------------------------------------------------------------------------------------------------------

type tokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Error        string `json:"error"`
	Message      string `json:"message"`
}

func (h *Handler) HandleAuthenticationTokenRequest(desmosAddress string, oAuthCode string) error {
	// Build the params values
	params := url.Values{}
	params.Add("client_id", h.cfg.ClientID)
	params.Add("client_secret", h.cfg.ClientSecret)
	params.Add("redirect_uri", h.cfg.RedirectURI)
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

	// Store the token inside the database
	oAuthToken := types.NewOAuthToken(desmosAddress, types.ServiceStreamlabs, resBody.AccessToken, resBody.RefreshToken)
	err = h.db.SaveOAuthToken(oAuthToken)
	if err != nil {
		return err
	}

	return err
}
