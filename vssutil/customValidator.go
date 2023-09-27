package vssutil

import (
	"regexp"

	"github.com/go-playground/validator"
)

const nonEnglishChars string = "àáâäæãåāăąçćčđďèéêëēėęěğǵḧîïíīįìłľḿñńǹňôöòóœøōõőṕŕřßśšşșťțûüùúūǘůűųẃẍÿýžźżÀÁÂÄÆÃÅĀĂĄÇĆČĐĎÈÉÊËĒĖĘĚĞǴḦÎÏÍĪĮÌŁĽḾÑŃǸŇÔÖÒÓŒØŌÕŐṔŔŘßŚŠŞȘŤȚÛÜÙÚŪǗŮŰŲẂẌŸÝŽŹŻ"
const (
	AlphaNumSpecialString = `^[a-zA-Z0-9\s` + nonEnglishChars + `!@#$_%^&*.,?()-=+:;|'<>-]+$`
	SimpleSentenceString  = `^[a-zA-Z0-9\s` + nonEnglishChars + `\.\,\!]+$`
	AlphaSpaceString      = `^[a-zA-Z0-9\s` + nonEnglishChars + `]+$`
)

var (
	AlphaNumSpecialRegex = regexp.MustCompile(AlphaNumSpecialString)
	SimpleSentenceRegex  = regexp.MustCompile(SimpleSentenceString)
	AlphaSpaceRegex      = regexp.MustCompile(AlphaSpaceString)
)

func AlphaNumSpecialValid(fl validator.FieldLevel) bool {
	return AlphaNumSpecialRegex.MatchString(fl.Field().String())
}
func SimpleSentenceValid(fl validator.FieldLevel) bool {
	return SimpleSentenceRegex.MatchString(fl.Field().String())
}
func AlphaSpaceValid(fl validator.FieldLevel) bool {
	return AlphaSpaceRegex.MatchString(fl.Field().String())
}
