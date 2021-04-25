package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}
