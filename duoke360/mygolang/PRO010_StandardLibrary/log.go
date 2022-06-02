package main

import (
	"fmt"
	"log"
	"os"
)

func print1(){
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, ",", age)
	// log.Fatal("致命错误！")
	fmt.Println("end...")
}

func panic1() {
	defer fmt.Println("panic 结束后再执行")
	log.Print("my log")
	log.Panic("my panic")
	log.Println("end...")
}

func setlog() {
	f, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(f)
	log.Print("my log...")
}

var logger *log.Logger

func init()  {
	logFile, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "success", log.Ldate | log.Ltime | log.Lshortfile)
}


func main() {
	//print1()
	//panic1()
	//setlog()
	logger.Println("自定义logger")

}
