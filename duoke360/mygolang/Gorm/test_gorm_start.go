package main
//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}
//
////创建表
//func create(db *gorm.DB) {
//	err := db.AutoMigrate(&Product{})
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
////插入数据
//func insert(db *gorm.DB) {
//	p := Product{
//		Code:  "M4A1",
//		Price: 3488,
//	}
//	db.Create(&p)
//}
//
////查询
//func find(db *gorm.DB) {
//	var product Product
//	//db.First(&product, 1)
//	//fmt.Printf("product:%v\n", product)
//
//	db.First(&product,"code=?", "M4A1")
//	fmt.Printf("product:%v\n", product)
//
//}
//
////更新
//func update(db *gorm.DB) {
//	var product Product
//	//db.First(&product,1)
//	//db.Model(&product).Update("Price",5000)
//	//
//	//db.Model(&product).Updates(Product{Code:"AK74M",Price: 6888})
//
//	db.First(&product,2)
//	db.Model(&product).Updates(map[string]interface{}{"Code":"M4A4","Price":8566})
//}
//
////删除
//func del(db *gorm.DB) {
//	var product Product
//	db.First(&product,1)
//	db.Delete(&product,1)
//}
//
//func main() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//update(db)
//	del(db)
//}
