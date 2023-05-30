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

type User struct {
	id int 
	loginname string
	password string
	username string
	state int
}

//这里定义了 多少个 就要查询 多少个  否则 无法 用限定个数的变量去接受
func query(){
	s := "SELECT * FROM user where id = ?"
	var u User
	err :=db.QueryRow(s,6).Scan(&u.id,&u.loginname,&u.password,&u.username,&u.state)
	if err != nil {
		fmt.Printf("数据查询错误%v\n", err)
	}else {
		fmt.Printf("查询到的用户信息%v\n", u)
	}
}

func queryMany(){
	s := "SELECT * FROM user"
	r,err :=db.Query(s)
	defer r.Close()
	var u User
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}else {
		for r.Next(){
			r.Scan(&u.id,&u.loginname,&u.password,&u.username,&u.state)
			fmt.Printf("查询到的用户数据%v\n", u)
		}
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
	//query()
	queryMany()

}