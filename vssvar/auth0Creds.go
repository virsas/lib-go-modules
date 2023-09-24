package vssvar

import (
	"os"

	"github.com/virsas/lib-go-modules/vsslib"
)

type Auth0Creds struct {
	Domain string
	Client string
	Secret string
}

func OPAuth0Creds(op vsslib.OpHandler, opAuth0Item string) (Auth0Creds, error) {
	var err error
	var auth0creds Auth0Creds = Auth0Creds{}

	auth0creds.Client, err = op.Get(opAuth0Item, "username")
	if err != nil {
		return auth0creds, err
	}

	auth0creds.Secret, err = op.Get(opAuth0Item, "password")
	if err != nil {
		return auth0creds, err
	}

	var auth0Domain string = "eu-west-1"
	auth0DomainValue, auth0DomainPresent := os.LookupEnv("AUTH0_DOMAIN")
	if auth0DomainPresent {
		auth0Domain = auth0DomainValue
	}
	auth0creds.Domain = auth0Domain

	return auth0creds, nil
}
