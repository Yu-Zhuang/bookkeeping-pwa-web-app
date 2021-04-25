package controller

import (
	"bookkeeping/config"
	"bookkeeping/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authCookie, err := c.Cookie(config.AuthCookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	mc, err := logic.ParseToken(authCookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	ok := logic.HasPersonByID(mc.Username)
	if ok == false {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	c.Set(config.AuthMidUserNameKey, mc.Username)
	c.Next()
}
