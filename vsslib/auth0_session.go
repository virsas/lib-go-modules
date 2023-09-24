package vsslib

import "gopkg.in/auth0.v5/management"

func NewAuth0Session(domain string, client string, secret string) (*management.Management, error) {
	var err error
	var sess *management.Management

	sess, err = management.New(domain, management.WithClientCredentials(client, secret))
	if err != nil {
		return nil, err
	}

	return sess, nil

}
