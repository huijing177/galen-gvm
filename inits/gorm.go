package inits

import (
	"fmt"
	"os"
	"strings"

	"galen-gvm/global"
	"galen-gvm/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 获取链接 默认是链接mysql
func Gorm() *gorm.DB {
	dbType := strings.ToLower(global.GVA_CONFIG.System.DbType)
	switch dbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgsql()
	default:
		return GormMysql()
	}
}

// 在数据库中创建表
func RegisterTables() {
	db := global.GVA_DB

	err := db.AutoMigrate(
		system.SysApi{},
		system.JwtBlacklist{},
		system.SysUser{},
	)
	if err != nil {
		global.GVA_LOG.Error("RegisterTables err:", zap.Error(err))
		os.Exit(0)
	}
	fmt.Println("register table success")

}
