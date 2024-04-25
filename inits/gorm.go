package inits

import (
	"fmt"
	"galen-gvm/global"
	"galen-gvm/model/system"
	"os"

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

// 在数据库中创建表
func RegisterTables() {
	db := global.GVA_DB

	err := db.AutoMigrate(
		system.SysApi{},
	)
	if err != nil {
		fmt.Println("TODO")
		os.Exit(0)
	}
	fmt.Println("register table success")

}
