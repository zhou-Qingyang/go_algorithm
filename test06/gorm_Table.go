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

func createTable(){
	db.AutoMigrate(&Baby{})
	fmt.Printf("Baby 表创建完成")

}

// 简单插入
func insertBaby(){

	var b Baby
	b = Baby{
		Name:"笑话",
		Age:20,
		BirthDay:time.Now(),
	}
	r :=db.Create(&b)
	fmt.Printf("新增返回的 影响行数%v\n", r.RowsAffected)
	fmt.Printf("新增ID :%v", b.ID)

}
//根据选择 字段插入
func insertBySelect(){
	b := Baby{
		Name:"woc",
		Age:10,
	}
	result := db.Select("Name","Age","CreatedAt").Create(&b)
	fmt.Printf("新增返回的 影响行数%v\n", result.RowsAffected)
	fmt.Printf("新增ID :%v", b.ID)
}
//根据多个实体类 添加
func batchInsert(){
	var babies = []Baby{{Name:"xiaoli",Age:99,},{Name:"xiaowang",Age:89}}
	result := db.Create(&babies)
	fmt.Printf("批量插入受影响的行数%v\n", result.RowsAffected)
}
//根据map集合 插入  不会执行生命周期的方法  并且默认字段  不会插入
func insertBabyByMap(){
	// db.Model(&Baby{}).Create(map[string]interface{}{
	// 	"Name":"map","Age":98,
	// })
	// fmt.Printf("根据map 集合插入\n")

	db.Model(&Baby{}).Create([]map[string]interface{}{
	{ "Name":"aaa","Age":11,},
	{ "Name":"222","Age":11,},
	})
	fmt.Printf("map多次插入")


}

//调用生命周期 初始化 
func (b *Baby)BeforeCreate(tx *gorm.DB)(err error){
	fmt.Println("BeforCreate.....")
	return 
}

type Baby struct {
	gorm.Model
	Name string 
	Age int
	BirthDay time.Time
}

func main() {
	//createTable()
	// insertBaby()
	//insertBySelect()

	//batchInsert()
	insertBabyByMap()

}