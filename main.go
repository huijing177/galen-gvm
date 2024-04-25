package main

import (
	"galen-gvm/global"
	"galen-gvm/inits"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	mainInit()
	if global.GVA_DB != nil {
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	time.Sleep(time.Second * 10)
}

func mainInit() {
	// 读取配置
	// 日志
	global.GVA_LOG = inits.Zap()
	// DB
	global.GVA_DB = inits.Gorm()

	if global.GVA_DB != nil {
		inits.RegisterTables()
	}
}
