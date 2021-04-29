package controller

import (
	"bookkeeping/config"
	"bookkeeping/logic"
	"bookkeeping/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var person model.Person
	c.Bind(&person)
	if logic.IsInputAccountOK(person) == false {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if logic.HasAccount(person) == false {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	tokenString, _ := logic.GenToken(person.ID, config.TokenExpireDuration)
	c.SetCookie(config.AuthCookieName, tokenString, config.AuthExpireDuration, "/", config.HostUrl, true, true)
	c.JSON(http.StatusOK, nil)
}

func LogOut(c *gin.Context) {
	c.SetCookie(config.AuthCookieName, "", -1, "/", config.HostUrl, false, true)
}

func Register(c *gin.Context) {
	var person model.Person
	c.Bind(&person)
	if logic.IsRegisterPersonOK(person) == false {
		fmt.Println("IsRegisterPersonOK == false")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "IsRegisterPersonOK == false",
		})
		return
	}
	if logic.CreatePerson(person) == false {
		fmt.Println("CreatePerson == false")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "CreatePerson == false",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GetUser(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"msg": mc.Username,
	})
	// c.Set(conf.AuthMidUserNameKey, mc.Username)
	// c.Next()
}
