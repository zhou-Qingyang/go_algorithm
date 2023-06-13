package main

import (
	"fmt"

	"github.com/zhou-Qingzhang/go-middleware/global"
	"github.com/zhou-Qingzhang/go-middleware/initalize"
	"github.com/zhou-Qingzhang/go-middleware/model/system"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = initalize.Viper()     // 初始化Viper
	global.GVA_LOG = initalize.Zap()      // 初始化zap日志库
	global.GVA_DB = initalize.GormMysql() // gorm连接数据库

	fmt.Println("Mysql Name is :", global.GVA_CONFIG.Mysql.Dbname)
	fmt.Println("Mysql Path is :", global.GVA_CONFIG.Mysql.Path)
	fmt.Println("Mysql Port is :", global.GVA_CONFIG.Mysql.Port)
	fmt.Println("Mysql Username is :", global.GVA_CONFIG.Mysql.Username)
	fmt.Println("Mysql Password is :", global.GVA_CONFIG.Mysql.Password)
	fmt.Println("Mysql Name is :", global.GVA_CONFIG.Mysql.Config)

	var reserList []system.Category
	err := global.GVA_DB.Model(system.Category{}).Find(&reserList)
	for _, v := range reserList {
		fmt.Printf("每个人的数据:", v)
		fmt.Println("")
	}
	global.GVA_LOG.Info("1.查询数据成功")

	// 记录日志
	global.GVA_LOG.Info("2. User logged in.", zap.String("username", "john"), zap.Any("metadata", map[string]string{"email": "john@example.com"}))
	global.GVA_LOG.Error("3. create directory", zap.Any("error:", err))
	global.GVA_LOG.Warn("4.日志警告", zap.Any("Warn", map[string]string{"Wran": "Warn Testing"}))
	global.GVA_LOG.Debug("5.Debug模式", zap.Any("Debug", map[string]string{"Wran": "Warn Testing"}))
}
