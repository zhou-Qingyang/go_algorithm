package global

import (
	"github.com/spf13/viper"
	"github.com/zhou-Qingzhang/go-middleware/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
)
