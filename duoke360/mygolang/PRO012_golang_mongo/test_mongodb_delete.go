package main
//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
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
//func del() {
//	ctx := context.TODO()
//	defer client.Disconnect(ctx)
//	collection := client.Database("go_db").Collection("Student")
//	dr, err := collection.DeleteMany(ctx, bson.D{{"age", 100}})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("dr.DeletedCount:%v\n", dr.DeletedCount)
//
//}
//
//func main() {
//	initDB()
//	del()
//
//}
