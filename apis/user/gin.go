package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/desmos-labs/plutus/types"
)

// DisconnectRequest represents the request body that must be used to disconnect a service
type DisconnectRequest struct {
	types.SignedRequest `json:",inline"`
	Platform            string `json:"platform"`
}

// RegisterHandlers registers all the handlers related to the user APIs
func RegisterHandlers(r *gin.Engine, handler *Handler) {
	r.Group("/user").
		GET("/:desmosAddress", func(c *gin.Context) {
			res, err := handler.GetUserData(c.Param("desmosAddress"))
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			c.JSON(http.StatusOK, res)
		}).
		DELETE("/integrations", func(c *gin.Context) {
			// Get the request body
			bodyBz, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			var request DisconnectRequest
			err = json.Unmarshal(bodyBz, &request)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			// Handle the request
			err = handler.HandleDisconnectRequest(request)
			if err != nil {
				apiutils.HandleError(c, err)
				return
			}

			c.Status(http.StatusOK)
		})
}
