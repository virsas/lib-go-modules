package vssutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func DebugClaims(claims jwt.MapClaims) {
	var logClaims bool = false
	logClaimsValue, logClaimsPresent := os.LookupEnv("CLAIMS_DEBUG")
	if logClaimsPresent {
		logClaimsValueBool, err := strconv.ParseBool(logClaimsValue)
		if err != nil {
			logClaims = false
		}
		logClaims = logClaimsValueBool
	}
	if logClaims {
		fmt.Println("######### Log Claims Function #########")
		fmt.Println(claims)
		fmt.Println("######### Log Claims Function #########")
	}
}
