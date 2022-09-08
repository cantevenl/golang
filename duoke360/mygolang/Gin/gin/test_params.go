package main

import "github.com/gin-gonic/gin"

func TestGet(c *gin.Context) {
	name := c.Query("username")
	pwd := c.DefaultQuery("password", "123")
	c.String(200, "username: %s, password: %s", name, pwd)
}

func TestPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "123")
	c.HTML(200, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})
}

func TestPathParam(c *gin.Context) {
	s := c.Param("name")
	s2 := c.Param("age")

	c.String(200, "Name:%s,Age:%s", s, s2)
}

func GoSearch(c *gin.Context) {
	c.HTML(200, "query.html", nil)
}

func search(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	key := c.PostForm("key")
	c.String(200, "page:%s,key:%s", page,key)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/testGet",TestGet)
	engine.POST("/testPost",TestPost)
	engine.GET("/testPath/:name/:age",TestPathParam)

	engine.GET("/goSearch",GoSearch)
	engine.POST("/search",search)
	engine.Run(":8888")
}