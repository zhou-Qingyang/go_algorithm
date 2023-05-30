package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

//连接数据库
func initDB()(err error){
	dsn := "root:123456@tcp(127.0.0.1)/zhou?charset=utf8mb4&parseTime=True"
	db,err := sql.Open("mysql",dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertUser(){
	sqlstr := "insert into user (loginname,username,password,state) values (?,?,?,?)"
	r,err := db.Exec(sqlstr,"admin","11111",999,1)
	if err != nil {
		fmt.Printf("数据插入错误%v\n", err)
		return 
	}else {
		i,errs := r.LastInsertId()
		if errs != nil {
			fmt.Printf("err:%v\n", errs)
		}
		fmt.Printf("插入生成的id:%v\n", i)
	}
}

func insert2(){
	r,err := db.Exec("INSERT INTO user (username,password) VALUES (?,?)","admin",123456)
	if err != nil {
		fmt.Printf("插入数据错误:%v\n", err)
	}else {
		id,_ := r.LastInsertId()
		fmt.Printf("插入的数据id:%v\n", id)
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

	insertUser()
}