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
	LowerDashString       = `^[a-z\-]+$`
	AlphaNumDashString    = `^[a-zA-Z0-9\-]+$`
)

var (
	AlphaNumSpecialRegex = regexp.MustCompile(AlphaNumSpecialString)
	SimpleSentenceRegex  = regexp.MustCompile(SimpleSentenceString)
	AlphaSpaceRegex      = regexp.MustCompile(AlphaSpaceString)
	LowerDashRegex       = regexp.MustCompile(LowerDashString)
	AlphaNumDashRegex    = regexp.MustCompile(AlphaNumDashString)
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
func LowerDashValid(fl validator.FieldLevel) bool {
	return LowerDashRegex.MatchString(fl.Field().String())
}
func AlphaNumDashValid(fl validator.FieldLevel) bool {
	return AlphaNumDashRegex.MatchString(fl.Field().String())
}
