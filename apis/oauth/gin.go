package oauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/desmos-labs/plutus/types"

	"github.com/gin-gonic/gin"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
)

// TokenRequest represents the request body that must be used to get an authorization token
type TokenRequest struct {
	types.SignedRequest `json:",inline"`
	Platform            string `json:"platform"`
	OAuthCode           string `json:"oauth_code"`
}

// RegisterHandlers registers all the handlers related to the Streamlabs APIs
func RegisterHandlers(r *gin.Engine, handler *Handler) {
	// Handle the requests
	r.Group("/oauth").
		POST("/token", func(c *gin.Context) {
			// Get the request body
			bodyBz, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			var request TokenRequest
			err = json.Unmarshal(bodyBz, &request)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			// Handle the request
			err = handler.HandleAuthenticationTokenRequest(request)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			c.Status(http.StatusOK)
		})
}
