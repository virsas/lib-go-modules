package router

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
)

func getCors() cors.Config {
	config := cors.DefaultConfig()

	var corsAllowedOrigin string = "*"
	corsAllowedOriginValue, corsAllowedOriginPresent := os.LookupEnv("CORS_ORIGIN")
	if corsAllowedOriginPresent {
		corsAllowedOrigin = corsAllowedOriginValue
	}

	config.AllowOrigins = strings.Split(strings.ReplaceAll(corsAllowedOrigin, " ", ""), ",")
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	config.MaxAge = 300

	return config
}
