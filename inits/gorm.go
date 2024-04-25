package inits

import (
	"galen-gvm/global"

	"gorm.io/gorm"
)

// 获取链接 默认是链接mysql
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgsql()
	default:
		return GormMysql()
	}
}
