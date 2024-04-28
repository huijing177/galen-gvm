package system

import (
	"galen-gvm/global"

	"go.uber.org/zap"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (JwtBlacklist) TableName() string {
	return "sys_jwt_blacklist"
}

func LoadAllJwtBlackList() ([]string, error) {
	var res []string

	err := global.GVA_DB.Model(&JwtBlacklist{}).Select("jwt").Find(&res).Error
	if err != nil {
		global.GVA_LOG.Error("LoadAllJwtBlackList err:", zap.Error(err))
	}
	return res, err
}
