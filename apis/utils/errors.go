package utils

import (
	"net/http"
	"unicode"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gin-gonic/gin"
)

func WrapErr(statusCode int, res string) error {
	return sdkerrors.New("http", uint32(statusCode), res)
}

func UnwrapErr(err error) (statusCode int, res string) {
	if sdkErr, ok := err.(sdkerrors.Error); ok {
		return int(sdkErr.ABCICode()), UcFirst(sdkErr.Error())
	}
	return http.StatusInternalServerError, UcFirst(err.Error())
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func HandleError(c *gin.Context, err error) {
	statusCode, res := UnwrapErr(err)
	c.String(statusCode, res)
}
