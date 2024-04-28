package inits

import (
	"galen-gvm/global"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func Cache() local_cache.Cache {
	expiresTime, err := cast.ToDurationE(global.GVA_CONFIG.Jwt.ExpiresTime)
	if err != nil {
		global.GVA_LOG.Error("Cache err:", zap.Error(err))
		expiresTime = time.Hour * 24 * 10
	}
	return local_cache.NewCache(
		local_cache.SetDefaultExpire(expiresTime),
	)
}
