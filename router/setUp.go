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
	r.GET("/zService-worker.js", func(c *gin.Context) {
		folder := "html"
		fileName := "zService-worker.js"
		path := "front/static/" + folder + "/" + fileName
		c.File(path)
	})
	r.GET("/", controller.HTMLhandler)
	r.GET("/:page", controller.HTMLhandler)

	api := r.Group("/api")
	{
		api.GET("/getUser", controller.GetUser)
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		api.POST("/logOut", controller.LogOut)

		api.POST("/addPayment", controller.Auth, controller.AddPayment)
		api.GET("/getPaymentHistory", controller.Auth, controller.GetPaymentHistory)
		api.GET("/getProfile", controller.Auth, controller.GetProfile)
		api.GET("/getChartData", controller.Auth, controller.GetChartData)
		api.GET("/getAverage", controller.Auth, controller.GetAverage)
		api.GET("/deletRecord/:_id", controller.Auth, controller.DeletRecord)
	}

	return r
}
