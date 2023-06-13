package initalize

import (
	"fmt"

	"github.com/zhou-Qingzhang/go-middleware/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err != nil {
		// 处理错误
		fmt.Println("数据库 初始化失败")
	}

	global.GVA_LOG.Info("数据库初始化成功!")
	return db
}
