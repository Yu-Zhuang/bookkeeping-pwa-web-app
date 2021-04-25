package controller

import (
	"bookkeeping/logic"
	"bookkeeping/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func SignUp(c *gin.Context) {
	var person model.Person
	c.Bind(person)
	if logic.IsInputPersonOK(person) == false {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"msg": "none",
	})
}
