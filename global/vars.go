package global

import (
	"galen-gvm/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	GVA_VP     *viper.Viper
)
