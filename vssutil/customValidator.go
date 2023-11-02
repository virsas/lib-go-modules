package vssutil

import (
	"regexp"

	"github.com/go-playground/validator"
)

const nonEnglishChars string = "àáâäæãåāăąçćčđďèéêëēėęěğǵḧîïíīįìłľḿñńǹňôöòóœøōõőṕŕřßśšşșťțûüùúūǘůűųẃẍÿýžźżÀÁÂÄÆÃÅĀĂĄÇĆČĐĎÈÉÊËĒĖĘĚĞǴḦÎÏÍĪĮÌŁĽḾÑŃǸŇÔÖÒÓŒØŌÕŐṔŔŘßŚŠŞȘŤȚÛÜÙÚŪǗŮŰŲẂẌŸÝŽŹŻ"
const allSpecialChars string = `'"#@%&_!,:;~=<>\.\-\+\*\?\$\^\(\)\[\]\{\}\/\\\|`

const (
	AlphaSpaceString      = `^[a-zA-Z` + nonEnglishChars + `\s]+$`
	AlphaNumSpaceString   = `^[a-zA-Z0-9` + nonEnglishChars + `\s]+$`
	AlphaNumSpecialString = `^[a-zA-Z0-9` + nonEnglishChars + allSpecialChars + `\s]+$`
	SimpleSentenceString  = `^[\w` + nonEnglishChars + `\s'"\-\.?!]+$`
	SlugString            = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
)

var (
	AlphaSpaceRegex      = regexp.MustCompile(AlphaSpaceString)
	AlphaNumSpaceRegex   = regexp.MustCompile(AlphaNumSpaceString)
	AlphaNumSpecialRegex = regexp.MustCompile(AlphaNumSpecialString)
	SimpleSentenceRegex  = regexp.MustCompile(SimpleSentenceString)
	SlugRegex            = regexp.MustCompile(SlugString)
)

func AlphaSpaceValid(fl validator.FieldLevel) bool {
	return AlphaSpaceRegex.MatchString(fl.Field().String())
}
func AlphaNumSpaceValid(fl validator.FieldLevel) bool {
	return AlphaNumSpaceRegex.MatchString(fl.Field().String())
}
func AlphaNumSpecialValid(fl validator.FieldLevel) bool {
	return AlphaNumSpecialRegex.MatchString(fl.Field().String())
}
func SimpleSentenceValid(fl validator.FieldLevel) bool {
	return SimpleSentenceRegex.MatchString(fl.Field().String())
}
func SlugValid(fl validator.FieldLevel) bool {
	return SlugRegex.MatchString(fl.Field().String())
}
