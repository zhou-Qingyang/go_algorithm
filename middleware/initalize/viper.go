package initalize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zhou-Qingzhang/go-middleware/global"
)

/**
viper.SetConfigFile("config.yaml")  // 设置配置文件名称和路径
viper.AddConfigPath(".")           // 添加搜索配置文件的路径
viper.ReadInConfig()                // 读取并解析配置文件
port := viper.GetInt("port")       // 获取配置项
*/

// Viper
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {

	v := viper.New()

	v.SetConfigFile("config.yaml") // 设置配置文件名称和路径
	v.AddConfigPath(".")           // 添加搜索配置文件的路径
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	//设置监听器
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
