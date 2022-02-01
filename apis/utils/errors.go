package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	c.AbortWithError(http.StatusInternalServerError, err)
}
