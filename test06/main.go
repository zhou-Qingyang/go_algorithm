package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func createTable(db *gorm.DB){
	//新建表
	db.AutoMigrate(&Product{})
}

func insertData(db *gorm.DB,pro Product){

	db.Create(&pro)
	fmt.Printf("数据插入完成\n")
}

func findData(db *gorm.DB){
	var p Product
	//根据主键查询
	// db.First(&p,5)
	// fmt.Printf("单行查询的用户信息%v\n", p)
	
	db.First(&p,"code = ? and price = ?","wowo",100)
	fmt.Printf("根据条件查询的用户信息%v\n", p)
}

func update(db *gorm.DB){
	var p Product
	db.First(&p,2)

	//第一种方式  	根据单个 字段修改
	//db.Model(&p).Update("PRICE",1000)

	//我靠  这里方法 调用是Updates

	//根据自身 的 实体修改
	// db.Model(&p).Updates(Product{
	// 	Code:"www",
	// 	Price:999999,
	// })

	//第三种方式  根据 map 修改
	db.Model(&p).Updates(map[string]interface{}{"Code":"nide","Price":1999})
	fmt.Printf("更新数据成功\n")
}

func deleteProduct(db *gorm.DB){
	var p Product
	db.First(&p,2)
	fmt.Printf("需要删除的数据%v\n", p)
	db.Delete(&p,2)
	fmt.Printf("删除数据成功")
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1)/zhou?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}

	// var pro Product
	// pro = Product{
	// 	Code:"wowo",
	// 	Price:100,
	// }

	// var pro = new(Product)
	// pro.Code = "nima"
	// pro.Price = 900
	// insertData(db,*pro)


	//findData(db)
	//update(db)
	deleteProduct(db)
}