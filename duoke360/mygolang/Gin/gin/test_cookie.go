package main

import "github.com/gin-gonic/gin"

func Handler1(c *gin.Context) {
	s, err := c.Cookie("username")
	if err != nil {
		s = "zhu"
		c.SetCookie("username", s, 60*60, "/", "localhost", false, true)
	}
	c.String(200,"测试cookie")
}

func main() {
	engine := gin.Default()
	engine.GET("/test_cookie", Handler1)
	engine.Run()
}
