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
	c.String(200, "form data : %s", user)
}

type User2 struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestGetBind(c *gin.Context) {
	var user2 User2
	err := c.ShouldBind(&user2)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User2:%s", user2)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/register", GoRegister)
	engine.POST("/register",Register1)
	//http://localhost:8888/testGetBind?username=zhu&password=12336456
	engine.GET("/testGetBind",TestGetBind)
	engine.Run(":8888")
}
