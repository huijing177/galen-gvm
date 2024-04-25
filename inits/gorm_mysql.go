package inits

import (
	"fmt"
	"galen-gvm/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	c := global.GVA_CONFIG.Mysql
	if c.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       c.Dsn(), // 配置信息
		DefaultStringSize:         191,     // 默认string长度
		SkipInitializeWithVersion: false,   // 是否根据版本配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,   // 表前缀
			SingularTable: c.Singular, // 是否开启全局禁用复数
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 迁移时，禁用外键
	})
	if err != nil {
		fmt.Println("TODO,可能是第一次没有链接信息,所以这里先不处理")
		global.GVA_LOG.Warn("TODO,可能是第一次没有链接信息,所以这里先不处理")
		return nil
	}
	db.InstanceSet("gorm:table_options", "ENGINE="+c.Engine)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	return db
}
