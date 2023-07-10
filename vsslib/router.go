package vsslib

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"go.uber.org/zap"
)

func NewRouter(assets string) (*gin.Engine, error) {
	logger, _ := zap.NewProduction()

	var debugLogging bool = false
	debugLoggingValue, debugLoggingPresent := os.LookupEnv("GIN_DEBUG")
	if debugLoggingPresent {
		debugLoggingValueBool, err := strconv.ParseBool(debugLoggingValue)
		if err == nil {
			debugLogging = debugLoggingValueBool
		}
	}

	if debugLogging {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	r := gin.Default()
	r.Use(cors.New(getCors()))
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(favicon.New(assets + "/favicon.ico"))

	return r, nil
}

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
