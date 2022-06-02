package main
//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//	"time"
//)
//
//var db *gorm.DB
//
//type User struct {
//	gorm.Model
//	Name     string
//	Age      int
//	Birthday time.Time `gorm:"2018-05-025 15:27:25"`
//}
//
//var user = User{
//	Name:"Zhu",
//	Age:4,
//	Birthday: time.Now(),
//}
//
//func init() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
//	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = database
//	fmt.Printf("db:%v\n", db)
//}
//
//func createTable() {
//	err := db.AutoMigrate(&User{})
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func insert() {
//	result := db.Create(&user)
//	fmt.Printf("result.RowsAffected:%v\n",result.RowsAffected)
//	fmt.Printf("user.ID:%v\n",user.ID)
//}
//
//func insert2() {
//	db.Select("Name","Age","CreatedAt").Create(&user)
//}
//
////批量插入
//func insertMany() {
//	var users = []User{{Name: "jinzhu1",Age:100}, {Name: "jinzhu2",Age:110}, {Name: "jinzhu3",Age:120}}
//	db.Create(&users)
//}
//
//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	fmt.Println("Before Create")
//	return
//}
//
//func insertMap() {
//	//db.Model(&User{}).Create(map[string]interface{}{
//	//	"Name": "xiaoniu", "Age": 18,
//	//})
//
//	db.Model(&User{}).Create([]map[string]interface{}{
//		{"Name": "zhu_1", "Age": 18},
//		{"Name": "zhu_2", "Age": 20},
//	})
//}
//
//func main() {
//	//createTable()
//	//insertMany()
//	insertMap()
//
//}
