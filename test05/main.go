package main
import (
	"time"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)
func main() {
	
	db,err := sql.Open("root:123456@/zhou")
	if err != nil {
		panic(err)
	}
	//最大连接时长
	db.SetConnMaxLifetime(time,Minute * 3)
	//最大连接数
	db.SetMaxOpenConnos(10)
	//空闲连接数
	db.SetMaxIdleConnos(10)
}