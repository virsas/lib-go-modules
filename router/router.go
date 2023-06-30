package router

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"go.uber.org/zap"
)

func New(assets string) (*gin.Engine, error) {
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
