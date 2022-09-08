package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/login",controller.GoLogin)
	e.POST("/login",controller.Login)

	//e.GET("/index", controller.ListUser)
	e.POST("/register", controller.Register)
	e.GET("/register",controller.GoRegister)

	e.GET("/",controller.Index)
	e.Run()
}
