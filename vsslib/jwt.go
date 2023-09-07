package vsslib

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

func JWTEncode(claims jwt.MapClaims, privateKey []byte) (string, error) {
	var err error
	var signedString string

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
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

func JWTDecode(token string, publicKey []byte) (jwt.MapClaims, error) {
	var err error

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return publicKey, nil })
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
