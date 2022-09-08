package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestMW(c *gin.Context) {
	c.String(200, "hello,%s", "zhu")
}

func MyMiddleware1(c *gin.Context) {
	fmt.Println("我的第一个中间件")
}

func MyMiddleware2(c *gin.Context) {
	fmt.Println("我的第二个中间件")
}

func main() {
	engine := gin.Default()
	engine.Use(MyMiddleware1, MyMiddleware2)
	engine.GET("/testmw", TestMW)
	engine.Run()
}
