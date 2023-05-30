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

func updateUser(){
	 s := "UPDATE user SET loginname = ? , password = ? where id = ?"
	 r,err := db.Exec(s,"zhouzju",8888,6)
	 if err != nil {
		fmt.Printf("数据库 修改错误%v\n", err)
		return
	 }else {
		i,_ := r.RowsAffected()
		fmt.Printf("修改受影响的行数:%v\n", i)
	 }
}

func deleteUser(){
	s := "DELETE FROM user WHERE id = ?"
	r,err := db.Exec(s,5)
	if err != nil {
		fmt.Printf("删除数据失败:%v\n",err)
		return
	}else {
		i,_ := r.RowsAffected()
		fmt.Printf("删除受影响的行数:%v\n", i)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("数据库连接失败!err:%v\n", err)
		return
	} else {
		fmt.Printf("数据库连接成功\n")
	}

	updateUser()
	deleteUser()
}