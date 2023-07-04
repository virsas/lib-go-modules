package utils

import (
	"regexp"

	"github.com/go-playground/validator"
)

const nonEnglishChars string = "àáâäæãåāăąçćčđďèéêëēėęěğǵḧîïíīįìłľḿñńǹňôöòóœøōõőṕŕřßśšşșťțûüùúūǘůűųẃẍÿýžźżÀÁÂÄÆÃÅĀĂĄÇĆČĐĎÈÉÊËĒĖĘĚĞǴḦÎÏÍĪĮÌŁĽḾÑŃǸŇÔÖÒÓŒØŌÕŐṔŔŘßŚŠŞȘŤȚÛÜÙÚŪǗŮŰŲẂẌŸÝŽŹŻ"
const (
	AlphaNumSpecialString = `^[a-zA-Z0-9\s` + nonEnglishChars + `!@#$_%^&*.,?()-=+:;|'<>-]+$`
)

var (
	AlphaNumSpecialRegex = regexp.MustCompile(AlphaNumSpecialString)
)

func AlphaNumSpecialValid(fl validator.FieldLevel) bool {
	return AlphaNumSpecialRegex.MatchString(fl.Field().String())
}
