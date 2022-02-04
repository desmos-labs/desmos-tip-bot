package oauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
)

// tokenRequestBody represents the request body that must be used to get an authorization token
type tokenRequestBody struct {
	Service       string `json:"service"`
	DesmosAddress string `json:"desmos_address"`
	OAuthCode     string `json:"oauth_code"`
}

// RegisterHandlers registers all the handlers related to the Streamlabs APIs
func RegisterHandlers(r *gin.Engine, handler *Handler) {
	// Handle the requests
	r.Group("/oauth").
		POST("/token", func(c *gin.Context) {
			// TODO: Authenticate the request maybe using a fake signed transaction

			// Get the request body
			bodyBz, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			var body tokenRequestBody
			err = json.Unmarshal(bodyBz, &body)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			// Handle the request
			err = handler.HandleAuthenticationTokenRequest(body.Service, body.DesmosAddress, body.OAuthCode)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			c.Status(http.StatusOK)
		})
}
