package vssutil

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetJWTToken(ctx *gin.Context) (string, error) {
	var token string

	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		return token, errors.New("authorization header issue")
	}

	token = strings.TrimSpace(splitToken[1])

	return token, nil
}
