package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "汤姆猫")
	if err != nil {
		fmt.Println("conn.Do err=", err)
	}

	//取出
	r, err := redis.String(conn.Do("Get", "name"))

	fmt.Println(r)

	//如果要从pool取出连接，一定保证连接池没有关闭
}
