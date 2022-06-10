package main

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.HTML(200, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/login", Login)
	engine.POST("/login",DoLogin)
	engine.Run(":8888")
}
