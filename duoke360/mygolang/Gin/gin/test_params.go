package main

import "github.com/gin-gonic/gin"

func testGet(c *gin.Context) {
	s := c.Query("username")
	pwd := c.DefaultQuery("password", "123")
	c.String(200, "username:%s, password:%s", s, pwd)
}

func testPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "888")

	c.HTML(200, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})
}

func testPathParam(c *gin.Context) {
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
	//engine.GET("/testGet",testGet)
	//engine.POST("/testPost",testPost)
	engine.GET("/testPath/:name/:age", testPathParam)
	engine.GET("/goSearch",GoSearch)
	engine.POST("/search",search)
	engine.Run(":8888")
}
