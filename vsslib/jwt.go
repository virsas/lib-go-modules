package vsslib

import (
	"errors"
	"os"
	"strings"

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

func JWTToken(authorizationToken string) (string, error) {
	var token string

	splitToken := strings.Split(authorizationToken, "Bearer")

	if len(splitToken) != 2 {
		return "", errors.New("authorization header issue")
	}

	token = strings.TrimSpace(splitToken[1])
	return token, nil
}

func JWTDecode(token string) (jwt.MapClaims, error) {
	var err error

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

	var pubKeyPath = rootPath + "keys/" + keyPrefix + "_jwtRS256.pub"

	verifyBytes, err := os.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return verifyKey, nil })
	if err != nil {
		validationError, _ := err.(*jwt.ValidationError)
		if validationError.Errors == jwt.ValidationErrorExpired {
			return nil, errors.New("expirationError")
		} else {
			return nil, err
		}
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("getClaimsError")
	}

	return claims, nil
}
