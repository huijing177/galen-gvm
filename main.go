package main

import (
	"galen-gvm/global"
	"galen-gvm/inits"
	"galen-gvm/internal"

	_ "go.uber.org/automaxprocs"
)

func main() {
	mainInit()
	if global.GVA_DB != nil {
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	internal.Run()
}

func mainInit() {
	// 读取配置
	global.GVA_VP = inits.Viper()
	// 日志
	global.GVA_LOG = inits.Zap()
	// DB
	global.GVA_DB = inits.Gorm()
	// redis
	if global.GVA_CONFIG.System.UseRedis {
		global.GVA_REDIS = inits.Redis()
	}
	// cache
	global.BlackCache = inits.Cache()

	if global.GVA_DB != nil {
		inits.RegisterTables()
		inits.LoadAllJwtBlackList2Cache()
	}
}
