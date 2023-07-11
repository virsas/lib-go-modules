package vssutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func DebugClaims(claims jwt.MapClaims) {
	var claimsDebug bool = false
	claimsDebugValue, claimsDebugPresent := os.LookupEnv("CLAIMS_DEBUG")
	if claimsDebugPresent {
		claimsDebugValueBool, err := strconv.ParseBool(claimsDebugValue)
		if err != nil {
			claimsDebug = false
		}
		claimsDebug = claimsDebugValueBool
	}
	if claimsDebug {
		debug(claims, "Log Claims Function")
	}
}

func DebugRequestBody(c *gin.Context) {
	var requestDebug bool = false
	requestDebugValue, requestDebugPresent := os.LookupEnv("REQUEST_DEBUG")
	if requestDebugPresent {
		requestDebugValueBool, err := strconv.ParseBool(requestDebugValue)
		if err != nil {
			requestDebug = false
		}
		requestDebug = requestDebugValueBool
	}
	if requestDebug {
		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

		body := map[string]interface{}{
			"Action":   c.Request.Method,
			"Resource": c.Request.RequestURI,
			"Body":     string(data),
		}

		debug(body, "Log Request Body")
	}
}

func debug(output interface{}, msg string) {
	var prettyDebug bool = true
	prettyDebugValue, prettyDebugPresent := os.LookupEnv("PRETTY_DEBUG")
	if prettyDebugPresent {
		prettyDebugValueBool, err := strconv.ParseBool(prettyDebugValue)
		if err != nil {
			prettyDebug = false
		}
		prettyDebug = prettyDebugValueBool
	}

	fmt.Println("######### ", msg, " #########")
	if prettyDebug {
		prettyOutput, err := json.MarshalIndent(output, "", "    ")
		if err != nil {
			fmt.Println(output)
		} else {
			fmt.Println(string(prettyOutput))
		}
	} else {
		fmt.Println(output)
	}
	fmt.Println("######### ", msg, " #########")
}
