package initalize

import (
	"fmt"

	"os"

	"github.com/zhou-Qingzhang/go-middleware/global"
	"github.com/zhou-Qingzhang/go-middleware/initalize/internal"
	"github.com/zhou-Qingzhang/go-middleware/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	//函数判断全局配置文件中指定的日志文件夹是否存在。
	//若该文件夹不存在，则会在控制台输出一条信息并通过 os.Mkdir 函数创建一个新的文件夹。
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	//第二行定义 cores 变量，并调用 internal.Zap.GetZapCores() 函数获取 Zap 日志核心。
	//该函数返回一个 []zapcore.Core 类型的切片，其中包含了多个不同级别（如 Debug、Info、Warn 等）的 Zap 日志核心。
	cores := internal.Zap.GetZapCores()

	//第三行调用 zap.New 函数，使用多个核心创建一个新的日志记录器。
	//该函数接收一个 zapcore.Core 类型的接口切片作为参数，使用 zapcore.NewTee 函数将所有核心合并成一个。
	logger = zap.New(zapcore.NewTee(cores...))

	//第五行使用 global.GVA_CONFIG.Zap.ShowLine 判断是否需要在日志中添加调用者的位置信息。
	//如果该值为 true，则通过 logger.WithOptions 函数添加一个 AddCaller 选项
	//使得日志记录器在输出日志时能够显示调用者的文件名和行号信息。
	if global.GVA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
