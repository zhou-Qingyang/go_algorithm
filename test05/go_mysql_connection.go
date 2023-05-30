package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var db *sql.DB

func initDB()(err error) {
	//定义一个初始化 数据库的函数
	dsn := "root:123456@tcp(127.0.0.1:3306)/zhou?charset=utf8mb4&parseTime=True"
	//这里 不能使用 := 我们是给全部变量赋值  然后zai main函数种使用全局变量
	db,err  = sql.Open("mysql",dsn)
	if err != nil {
		return err
	}
	//这里开始连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}

func main() {
	
	// db,err := sql.Open("mysql","root:123456@/zhou")
	// if err != nil {
	// 	panic(err)
	// }
	// //最大连接时长
	// db.SetConnMaxLifetime(time.Minute*3)
	// //最大连接数
	// db.SetMaxOpenConns(10)
	// //空闲连接数
	// db.SetMaxIdleConns(10)

	// fmt.Printf("数据库连接%v", db)
	err := initDB()
	if err != nil {
		fmt.Printf("数据库连接失败!err:%v\n", err)
		return
	} else {
		fmt.Printf("数据库连接成功\n")
	}
}