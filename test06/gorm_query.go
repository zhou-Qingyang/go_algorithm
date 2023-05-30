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

//单条记录查询
func selectOne(){
	var b Baby
	//SELECT * FROM order by id LIMIT 1
	db.First(&b)
	
	// db.Take(&b)
	// db.Last(&b)
	fmt.Printf("查询的记录:%v\n",b)
}

func selectMore(){
	var babys []Baby
	db.Find(&babys,[]int{1,2,3})
	// for range 匿名变量  最好 还是 _,i 否则默认是 下表
	for  _,baby := range babys {
		fmt.Printf("查询的宝贝信息 : %v\n", baby)
	}

}

func selectAll(){
	var babys []Baby
	result := db.Find(&babys)
	for _,v := range babys {
		fmt.Printf("插叙:%v\n", v)
		
	}
	fmt.Printf("查询的总条数%v\n", result.RowsAffected)
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
	// selectOne()

	//selectMore()

	selectAll()


}