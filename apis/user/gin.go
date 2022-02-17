package user

import (
	apiutils "github.com/desmos-labs/plutus/apis/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		})
}
