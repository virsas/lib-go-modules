package vssutil

import "regexp"

const (
	updatedString       = `,"createdAt":\s?"(20.*?)"`
	createdString       = `,"updatedAt":\s?"(20.*?)"`
	authorizationString = `"authorization":\s?"(.*?)"`
)

var (
	UpdatedRegex       = regexp.MustCompile(updatedString)
	CreatedRegex       = regexp.MustCompile(createdString)
	AuthorizationRegex = regexp.MustCompile(authorizationString)
)
