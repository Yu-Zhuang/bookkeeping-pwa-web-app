package controller

import (
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
	c.JSON(http.StatusOK, nil)
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
	c.JSON(http.StatusUnauthorized, gin.H{
		"msg": "none",
	})
}
