package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go 向redis写入数据和读取数据
	//1.连接到redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis连接失败，错误=", err)
		return
	}
	//关闭连接
	defer conn.Close()

	//2.通过go向redis写入数据 string[key-val]
	_, err = conn.Do("Set", "name", "zhuzhu")
	if err != nil {
		fmt.Println("set错误=", err)
		return
	}

	//3.通过go向redis读取数据 string[key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get错误=", err)
		return
	}

	//因为返回的r是一个interface{}
	//因为name对应的值是string，因此我们需要转换
	fmt.Println("操作成功 值是=", r)

}
