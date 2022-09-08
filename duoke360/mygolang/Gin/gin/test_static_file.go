package main

import (
	"github.com/gin-gonic/gin"
	"log"
	//"net/http"
)

func GoStatic(c *gin.Context) {
	c.HTML(200,"test_static.html",nil)
}

func main() {
	engine := gin.Default()
	engine.Static("/assets","./assets")
	//engine.StaticFS("/droot",http.Dir("d:/"))
	engine.StaticFile("/favicon.ico", "./assets/favicon.ico")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/go",GoStatic)

	err := engine.Run()
	if err != nil {
		log.Fatal(err)
	}
}
