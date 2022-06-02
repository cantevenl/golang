package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time `gorm:"2018-05-025 15:27:25"`
}


func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = database
	fmt.Printf("db:%v\n", db)
}

func select1() {
	var user User
	//db.First(&user,17)
	//fmt.Printf("user.ID:%v\n", user.ID)
	////fmt.Printf("user:%v\n", user)
	//db.Take(&user)
	//fmt.Printf("user.ID:%v\n", user.ID)
	db.Last(&user)
	fmt.Printf("user.ID:%v\n", user.ID)

}

func select2() {
	var user User
	db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;
	fmt.Printf("user.ID:%v\n", user.ID)

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;
	fmt.Printf("user.ID:%v\n", user.ID)

	var users []User
	db.Find(&users, []int{1,2,3})
	// SELECT * FROM users WHERE id IN (1,2,3);
	for _,user := range users {
		fmt.Printf("user.ID:%v\n", user.ID)

	}
}

func select3() {
	var users []User
	// 获取全部记录
	result := db.Find(&users)
	// SELECT * FROM users;
	fmt.Println(result.RowsAffected)
	//result.RowsAffected // 返回找到的录数，相当于 `len(users)`
	//result.Error        // returns error记


}


func main() {
	select3()

}
