package main

import "github.com/gin-gonic/gin"

func goRegister(c *gin.Context) {
	c.HTML(200,"form.html",nil)
}


func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	hobby := c.PostFormArray("hobby")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	c.String(200, "Username:%s, Password:%s, hobby:%s, gender:%s, city:%s", username, password, hobby, gender, city)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/testForm",goRegister)
	engine.POST("/register",Register)
	engine.Run(":8888")
}
