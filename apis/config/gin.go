package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandlers(r *gin.Engine, handler *Handler) {
	r.GET("/config", func(c *gin.Context) {
		config := handler.HandleRequestConfig()
		c.JSON(http.StatusOK, config)
	})
}
