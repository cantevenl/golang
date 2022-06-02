package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB
//
//func initDB() (err error) {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
//	// 不会校验账号密码是否正确
//	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	//尝试与数据库建立连接（校验dsn是否正确）
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//
//	return nil
//
//}

func main() {
	//db, err := sql.Open("mysql", "root:123456@/go_db")
	//if err != nil {
	//	panic(err)
	//}
	//db.SetConnMaxLifetime(time.Minute * 3)
	//// 最大连接数
	//db.SetMaxOpenConns(10)
	//// 空闲连接数
	//db.SetMaxIdleConns(10)
	//
	//fmt.Printf("db: %v\n", db)
	//
	//err := initDB() // 调用输出化数据库的函数
	//if err != nil {
	//	fmt.Printf("初始化失败！,err:%v\n", err)
	//	return
	//}else{
	//	fmt.Printf("初始化成功")
	//}

}
