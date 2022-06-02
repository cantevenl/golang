package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//// 定义一个全局对象db
//var db *sql.DB
//
//// 定义一个初始化数据库的函数
//func initDB() (err error) {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
//	// 不会校验账号密码是否正确
//	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	// 尝试与数据库建立连接（校验dsn是否正确）
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//type User struct {
//	id       int
//	username string
//	password string
//}
//
//func queryOneRow() {
//	var u User
//	s := "select * from user_tbl where id =?"
//	err := db.QueryRow(s, 2).Scan(&u.id, &u.username, &u.password)
//	if err != nil {
//		fmt.Printf("err %v\n", err)
//	} else {
//		fmt.Printf("u:%v\n", u)
//	}
//}
//
//func queryManyRow() {
//	var u User
//	s := "select * from user_tbl"
//	r, err := db.Query(s)
//	defer r.Close()
//	if err != nil {
//		fmt.Printf("err %v\n", err)
//	} else {
//		for r.Next() {
//			r.Scan(&u.id, &u.username, &u.password)
//			fmt.Printf("id:%d username:%s password:%s\n", u.id, u.username, u.password)
//		}
//	}
//}
//
//func main() {
//	err := initDB() // 调用输出化数据库的函数
//	if err != nil {
//		fmt.Printf("初始化失败！,err:%v\n", err)
//		return
//	} else {
//		fmt.Printf("初始化成功\n")
//	}
//
//	//queryOneRow()
//	queryManyRow()
//}
