package controller

import (
	"bookkeeping/logic"
	"bookkeeping/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var person model.Person
	c.Bind(person)
	if logic.IsInputAccountOK(person) == false {
		c.JSON(http.StatusBadRequest, nil)
	}
	if logic.HasAccount(person) == false {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, nil)
}

func Register(c *gin.Context) {
	var person model.Person
	c.Bind(person)
	if logic.IsRegisterPersonOK(person) == false {
		c.JSON(http.StatusBadRequest, nil)
	}
	if logic.CreatePerson(person) == false {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, nil)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"msg": "none",
	})
}
