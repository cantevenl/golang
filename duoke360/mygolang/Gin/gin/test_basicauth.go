package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

//模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func Handler(c *gin.Context) {
	// 获取用户，它是由 BasicAuth 中间件设置的
	user := c.MustGet(gin.AuthUserKey).(string)
	fmt.Println(user)
	//if secret, ok := secrets[user]; ok {
	//	c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	//} else {
	//	c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	//}
}

func main (){
	engine := gin.Default()
	authorized := engine.Group("/admin",gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets",Handler)

	engine.Run()
}