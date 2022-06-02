package main
//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//)
//
//var client *mongo.Client
//var err error
//
//type Student struct {
//	Name string
//	Age  int
//}
//
//func initDB() {
//	//设置客户端连接配置
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//	//连接到Mongodb
//	client, err = mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//	//检查连接
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		fmt.Println("连接成功")
//	}
//}
//
//func insert() {
//	s1 := Student{
//		"tom",
//		20,
//	}
//
//	collection := client.Database("go_db").Collection("Student")
//	one, err := collection.InsertOne(context.TODO(), s1)
//	if err != nil {
//		fmt.Printf("err:%v\n", err)
//	} else {
//		fmt.Printf("one.InsertedID:%v\n", one.InsertedID)
//	}
//
//}
//
//func insertMany() {
//	s1 := Student{
//		"zhu",
//		2,
//	}
//
//	s2 := Student{
//		"mao",
//		1,
//	}
//
//	stus := []interface{}{s1, s2}
//
//	collection := client.Database("go_db").Collection("Student")
//	imr, err := collection.InsertMany(context.TODO(), stus)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("imr.InsertedIDs:%v\n", imr.InsertedIDs)
//}
//
//func main() {
//	initDB()
//	//insert()
//	insertMany()
//}
