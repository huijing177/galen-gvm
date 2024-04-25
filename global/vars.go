package global

import (
	"galen-gvm/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
)
