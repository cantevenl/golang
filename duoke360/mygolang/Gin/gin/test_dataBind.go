package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Username string   `form:"username"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func Register1(c *gin.Context) {
	var user User
	c.ShouldBind(&user)

	c.String(200, "form data:%s", user)
}

type User2 struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestGetBind(c *gin.Context) {
	var user User2
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}

type User3 struct {
	Username string `uri:"username"`
	Password string `uri:"password"`
}

func TestUrlBind(c *gin.Context) {
	var user User3
	err := c.ShouldBindUri(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/register", GoRegister)
	engine.POST("/register", Register1)

	// http://localhost:8888/testGetBind?username=zhu&password=123
	engine.GET("/testGetBind", TestGetBind)
	// http://localhost:8888/testUrlBind/niu/123
	engine.GET("/testUrlBind/:username/:password",TestUrlBind)
	engine.Run(":8888")
}
