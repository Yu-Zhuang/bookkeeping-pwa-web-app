package logic

import (
	"bookkeeping/config"
	"bookkeeping/model"
	"fmt"
	"unicode"
)

func IsInputAccountOK(p model.Person) bool {
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

func IsRegisterPersonOK(p model.Person) bool {
	if p.ID == "" || p.Name == "" || p.Email == "" || p.Password == "" {
		fmt.Println("empty person")
		return false
	}
	IDlen, PWDlen := len(p.ID), len(p.Password)
	if IDlen < config.AccountMinLen || IDlen > config.AccountMaxLen || PWDlen < config.AccountMinLen || PWDlen > config.AccountMaxLen {
		fmt.Println("account length out of range")
		return false
	}
	if StringOnlyHasDigitAlpha(p.ID) && StringOnlyHasDigitAlpha(p.Password) {
		return true
	}
	fmt.Println("account has word which not belong number and alphabat")
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
