package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, handler *Handler) {
	r.GET("/config", func(c *gin.Context) {
		config := handler.HandleRequestConfig()
		c.JSON(http.StatusOK, config)
	})
}
