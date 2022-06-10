package main

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.String(200,"hello golang,%s","niu")
}

func Hi(c *gin.Context) {
	c.JSON(200,gin.H{
		"name":"猪",
		"age":"2",
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/hello",Hi)
	//默认监听8080
	engine.Run(":8989")
}