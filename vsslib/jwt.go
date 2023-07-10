package vsslib

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func JWTEncode(claims jwt.MapClaims) (string, error) {
	var err error
	var signedString string

	var rootPath string = "./"
	rootPathValue, rootPathPresent := os.LookupEnv("ROOT_PATH")
	if rootPathPresent {
		rootPath = rootPathValue
	}

	var keyPrefix string = "production"
	keyPrefixValue, keyPrefixPresent := os.LookupEnv("KEY_PREFIX")
	if keyPrefixPresent {
		keyPrefix = keyPrefixValue
	}

	var privKeyPath = rootPath + "keys/" + keyPrefix + "_jwtRS256.key"

	signBytes, err := os.ReadFile(privKeyPath)
	if err != nil {
		return signedString, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return signedString, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, err = token.SignedString(signKey)
	if err != nil {
		return signedString, err
	}

	return signedString, nil
}
