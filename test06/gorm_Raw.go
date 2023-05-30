package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"
)

var db *gorm.DB

func init() {
	dsn :="root:123456@tcp(127.0.0.1)/zhou?charset=utf8mb4&parseTime=True&loc=Local"
	r,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic(" Fatal to connect MYSQL....")
	}
	db = r
}


func selectByRow(){
	var baby Baby
	db.Raw("Select b.name,b.id,b.age from babies b where name = ?","nong").Scan(&baby)
	fmt.Printf("原生sql 查询的:%v\n", baby)
}

func selectCount(){
	var co int
	db.Raw("SELECT count(*) from babies b ").Scan(&co)
	fmt.Printf("总记录条数:%v\n", co)
}

func updateBaby(){
	db.Exec("update babies set age = ? where name = ?",18,"nong")
	fmt.Printf("更新操作完成")
}

func SelectByNameArgs(){

	// //类似于 where name1 = "jizhou" or name2 = "jizhou"
	// db.Where("name1 = @name OR name2 = @name",sql.Named("name","jizhou")).Find(&baby)

	// //where name1 = "jinzhou2" or name2 = "jinzhou2"
	// db.Where("name1 = @name OR name2 = @name",map[string]interface{}{
	// 	"name":"jinzhou2",
	// }).Find(&baby)

	// //类似于 where name1 = "jizhou" or name2 = "jizhou" or name3 = jinzhou3
	// db.Where("name1 = @name1 OR name2 = @name1 OR name3 = @name2",sql.Named("name1","jizhou"),sql.Named("name2","jinzhou2")).Find(&baby)


}

func selectNameAndAge(){
	var name string
	var age int
	row := db.Table("babies").Where("id = ?",5).Select("name","age").Row()
	row.Scan(&name,&age)
	fmt.Printf("查询名字:%v\n", name)
	fmt.Printf("查询年龄:%v\n", age)
}

func selectRowByOnly(){
	var name string
	var age int

	rows,_ := db.Model(&Baby{}).Where("id > ?",5).Select("name","age").Rows()
	for rows.Next(){
		rows.Scan(&name,&age)
		fmt.Printf("name:%v\n", name)
		fmt.Printf("age:%v\n", age)
		fmt.Printf("============\n")
	}
}


type Baby struct {
	gorm.Model
	Name string 
	Age int
	BirthDay time.Time
}

func main() {

	// selectByRow()
	// selectCount()
	// updateBaby()
	// selectNameAndAge()
	selectRowByOnly()

}