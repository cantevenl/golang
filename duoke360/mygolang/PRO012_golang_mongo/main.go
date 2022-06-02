package main
//
//import (
//	"context"
//	"fmt"
//	_ "go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//)
//
//var client *mongo.Client
//var err error
//
//func initDB() {
//	//设置客户端连接配置
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//	//连接到Mongodb
//	client, err = mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}else {
//		fmt.Printf("client:%v\n",client)
//	}
//	//检查连接
//	err = client.Ping(context.TODO(), nil)
//	if err !=nil {
//		log.Fatal(err)
//	}else {
//		fmt.Println("连接成功")
//	}
//}
//
//
//func main() {
//	initDB()
//}
