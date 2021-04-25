package config

import "time"

const (
	HostUrl       = "https://bookkepping.herokuapp.com/"
	AccountMinLen = 8
	AccountMaxLen = 20

	AuthExpireDuration  = 60 * 60 // unit: second
	AuthCookieName      = "_auth_token"
	TokenExpireDuration = time.Hour
)

var (
	TokenMySecret = []byte("fdspaij23890rhuisdahfjl8eFEWisdap9Rwqe")
)
