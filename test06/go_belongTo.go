package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

var db *gorm.DB

func init() {
	dsn :="root:123456@tcp(127.0.0.1)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	r,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic(" Fatal to connect MYSQL....")
	}
	db = r
}

//这里需要一个 外键 连接 两张表之间的关系
type User struct {
	ID int
	Name string
	Age int
	CompanyID int
	Company Company
}

type Company struct {
	ID int
	Name string
}

//一对一  每张表的主键 都可以作为 另一张表的外键
type Person struct {
	ID int
	Name string
	CreaditCard CreaditCard
}

type CreaditCard struct {
	ID int 
	PersonID uint
	Number int
}


// 一对多  一个学生 有多本书
type Student struct {
	ID int
	Name string
	Book []Book
}

type Book struct {
	ID int
	Name string
	StudentID uint
}

// 多对多  一个学生 有多本书
type Teacher struct {
	ID int
	Name string
	Languages []Language `gorm:"many2many:teacher_languages"`
}

type Language struct {
	ID int
	Name string
	
}
//创建表 父类在前
func createTableBelong(){
	db.AutoMigrate(&User{},&Company{})
	fmt.Printf("初始化表结构成共")
}

func createByOne(){
	db.AutoMigrate(&Person{},&CreaditCard{})
	fmt.Printf("新建一对一的关联关系")
}

func createOneToMany(){
	db.AutoMigrate(&Student{},&Book{})
	fmt.Printf("一对多表单关系")
}

func createManyToMany(){
	db.AutoMigrate(&Teacher{},&Language{})
	fmt.Printf("多对多表单关系")
}

func main() {
	
	//createTableBelong()
	//createByOne()
	createOneToMany()
}