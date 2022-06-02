package main
//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//	"time"
//)
//
//type Student struct {
//	Name string
//	Age  int
//	Addr string
//}
//
//var client *mongo.Client
//
//func initDB() {
//	// 连接到MongoDB
//	var err error
//	client, err = mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 检查连接
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connected to MongoDB!\n-------------------------------------------")
//}
//
////func insertOne() {
////	initDB()
////	s := Student{
////		Name: "jack",
////		Age:  40,
////		Addr: "NingXia",
////	}
////	collection := client.Database("student").Collection("student")
////	insertResult, err := collection.InsertOne(context.TODO(), s)
////	if err != nil {
////		log.Fatal(err)
////	}
////	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
////}
//
//func insertOne(s Student) {
//	initDB()
//	collection := client.Database("face_center").Collection("student")
//	insertResult, err := collection.InsertOne(context.TODO(), s)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
//}
//
//func insertMore(student []interface{}) {
//	initDB()
//	collection := client.Database("face_center").Collection("student")
//	insertManyResult, err := collection.InsertMany(context.TODO(), student)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
//}
//
//func find() {
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	collection := client.Database("face_center").Collection("student")
//	cur, err := collection.Find(ctx, bson.D{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer cur.Close(ctx)
//	for cur.Next(ctx) {
//		var result bson.D
//		err := cur.Decode(&result)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("result: %v\n", result)
//		fmt.Printf("result.Map(): %v\n", result.Map()["name"])
//	}
//	if err := cur.Err(); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func main() {
//	initDB()
//	find()
//	s := Student{
//		"tom",
//		28,
//		"tianjin",
//	}
//	insertOne(s)
//	//insertOne()
//	find()
//
//	s1 := Student{
//		"jack",
//		20,
//		"taiwan",
//	}
//	s2 := Student{
//		"kite",
//		21,
//		"aomen",
//	}
//	s3 := Student{
//		"rose",
//		23,
//		"nanjing",
//	}
//
//	student := []interface{}{s1, s2, s3}
//	insertMore(student)
//
//	find()
//
//}
