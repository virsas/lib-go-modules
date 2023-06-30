package router

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
)

func getCors() cors.Config {
	config := cors.DefaultConfig()

	var originsAllowed string = "*"
	originsAllowedValue, originsAllowedPresent := os.LookupEnv("CORS_ORIGINS")
	if originsAllowedPresent {
		originsAllowed = originsAllowedValue
	}

	var methodsAllowed string = "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	methodsAllowedValue, methodsAllowedPresent := os.LookupEnv("CORS_METHODS")
	if methodsAllowedPresent {
		methodsAllowed = methodsAllowedValue
	}

	var headersAllowed string = "Authorization, Content-Type"
	headersAllowedValue, headersAllowedPresent := os.LookupEnv("CORS_HEADERS")
	if headersAllowedPresent {
		headersAllowed = headersAllowedValue
	}

	var credentialsAllowed bool = true
	credentialsAllowedValue, credentialsAllowedPresent := os.LookupEnv("CORS_CREDENTIALS")
	if credentialsAllowedPresent {
		credentialsAllowedValueBool, err := strconv.ParseBool(credentialsAllowedValue)
		if err == nil {
			credentialsAllowed = credentialsAllowedValueBool
		}
	}

	var maxAge time.Duration = 300
	maxAgeValue, maxAgePresent := os.LookupEnv("CORS_MAXAGE")
	if maxAgePresent {
		maxAgeValueInt, err := strconv.Atoi(maxAgeValue)
		if err == nil {
			maxAge = time.Duration(maxAgeValueInt)
		}
	}

	config.AllowOrigins = strings.Split(strings.ReplaceAll(originsAllowed, " ", ""), ",")
	config.AllowMethods = strings.Split(strings.ReplaceAll(methodsAllowed, " ", ""), ",")
	config.AllowHeaders = strings.Split(strings.ReplaceAll(headersAllowed, " ", ""), ",")
	config.AllowCredentials = credentialsAllowed
	config.MaxAge = maxAge

	return config
}
