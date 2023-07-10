package vssutil

import "regexp"

const (
	authorizationString = `"authorization":\s?"(.*?)"`
)

var (
	AuthorizationRegex = regexp.MustCompile(authorizationString)
)
