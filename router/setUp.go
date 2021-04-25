package router

import (
	"bookkeeping/controller"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("front/html/*")
	r.GET("/static/:folder/:file", controller.Static)
	r.GET("/favicon.ico", controller.Favicon)
	r.GET("/", controller.HTMLhandler)
	r.GET("/:page", controller.HTMLhandler)

	api := r.Group("/api")
	{
		api.GET("/getUser", controller.GetUser)
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		api.POST("/addPayment", controller.AddPayment)
	}

	return r
}
