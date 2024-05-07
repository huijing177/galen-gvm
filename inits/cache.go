package inits

import (
	"galen-gvm/global"
	"galen-gvm/utils"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

func Cache() local_cache.Cache {
	expiresTime, err := utils.ParseDuration(global.GVA_CONFIG.Jwt.ExpiresTime)
	if err != nil {
		global.GVA_LOG.Error("Cache err:", zap.Error(err))
		global.GVA_LOG.Error("we will use default Jwt expires_time: 7day")
		expiresTime = time.Hour * 24 * 7
	}
	return local_cache.NewCache(
		local_cache.SetDefaultExpire(expiresTime),
	)
}
