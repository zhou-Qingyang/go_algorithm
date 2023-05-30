package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"

)
var db *gorm.DB

//Logger:logger.Default.Model(logger.info)
func init() {
	dsn :="root:123456@tcp(127.0.0.1)/zhou?charset=utf8mb4&parseTime=True&loc=Local"
	r,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic(" Fatal to connect MYSQL....")
	}
	db = r
}

type Baby struct {
	gorm.Model
	Name string
	BirthDay time.Time
	Age int	
}

//先 查询出来  再进行修改
func updateByOne(){
	var baby Baby
	db.First(&baby)  //id == 1
	baby.Name = "big tom"
	baby.Age = 88
	db.Save(&baby)

	fmt.Printf("修改用户信息\n")
}

func updateByNoSelect(){
	var baby Baby
	db.Model(&baby).Where("Age = ?",88).Update("name","nong")
	fmt.Printf("不查询修改信息成功\n")
}

func main() {
	// updateByOne()
	updateByNoSelect()
}