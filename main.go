package main

import (
	_ "galen-gvm/docs"
	"galen-gvm/global"
	"galen-gvm/inits"
	"galen-gvm/internal"

	_ "go.uber.org/automaxprocs"
)

// @title           GALEN-GVM Swagger API
// @version         1.0
// @description     This is a sample galen's server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8888
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
