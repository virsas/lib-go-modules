package vsslib

import (
	"crypto/rsa"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/golang-jwt/jwt"
)

func RSAPublicFile() (*rsa.PublicKey, error) {
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

	var keyPath = rootPath + "keys/" + keyPrefix + "_jwtRS256.pub"

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	verifiedPem, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return verifiedPem, nil
}

func RSAPrivateFile() (*rsa.PrivateKey, error) {
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

	var keyPath = rootPath + "keys/" + keyPrefix + "_jwtRS256.key"

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	verifiedPem, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return verifiedPem, nil
}

func RSAPublicUrl(pemURL string) (*rsa.PublicKey, error) {
	keyURL, err := url.Parse(pemURL)
	if err != nil {
		return nil, err
	}

	keyRESP, err := http.Get(keyURL.String())
	if err != nil {
		return nil, err
	}
	defer keyRESP.Body.Close()

	if keyRESP.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error status code: %d", keyRESP.StatusCode)
	}

	keyBytes, err := io.ReadAll(keyRESP.Body)
	if err != nil {
		return nil, err
	}

	verifiedPem, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return verifiedPem, nil
}
