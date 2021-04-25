package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTMLhandler(c *gin.Context) {
	page := c.Param("page")
	if page == "" {
		page = "index"
	}
	page = page + ".html"
	fmt.Println(page)
	c.HTML(http.StatusOK, page, nil)
}

func Static(c *gin.Context) {
	folder := c.Param("folder")
	fileName := c.Param("file")
	path := "front/static/" + folder + "/" + fileName
	c.File(path)
}
func Favicon(c *gin.Context) {
	folder := "image"
	fileName := "favicon.ico"
	path := "front/static/" + folder + "/" + fileName
	c.File(path)
}
