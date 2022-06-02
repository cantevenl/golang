package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client
var err error

type Student struct {
	Name string
	Age  int
}

func initDB() {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接到Mongodb
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("连接成功")
	}
}

func find() {
	ctx := context.TODO()
	defer client.Disconnect(ctx)
	collection := client.Database("go_db").Collection("Student")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.D
		cursor.Decode(&result)

		fmt.Printf("result:%v\n",result)
		fmt.Printf("name:%v,age:%v\n",result.Map()["name"],result.Map()["age"])
	}
}

func main() {
	initDB()
	find()

}
