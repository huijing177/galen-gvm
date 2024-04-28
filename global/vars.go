package global

import (
	"galen-gvm/config"

	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	GVA_VP     *viper.Viper
	GVA_REDIS  *redis.Client
	BlackCache local_cache.Cache
)
