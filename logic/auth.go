package logic

import (
	"bookkeeping/config"
	"bookkeeping/model"
	"unicode"
)

func IsInputPersonOK(p model.Person) bool {
	IDlen := len(p.ID)
	PWDlen := len(p.Password)
	if IDlen < config.AccountMinLen || IDlen > config.AccountMaxLen || PWDlen < config.AccountMinLen || PWDlen > config.AccountMaxLen {
		return false
	}
	if StringOnlyHasDigitAlpha(p.ID) && StringOnlyHasDigitAlpha(p.Password) {
		return true
	}
	return false
}

func StringOnlyHasDigitAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			continue
		}
		return false
	}
	return true
}
