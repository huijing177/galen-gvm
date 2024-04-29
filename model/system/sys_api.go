package system

import "galen-gvm/global"

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"column:path;comment:api路径"` // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`          // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (SysApi) TableName() string {
	return "sys_apis"
}

func GetAllApis() ([]SysApi, error) {
	var apis []SysApi
	err := global.GVA_DB.Find(&apis).Error
	return apis, err
}

func GetApiById(id int) (SysApi, error) {
	var res SysApi
	err := global.GVA_DB.First(&res, "id = ?", id).Error
	return res, err
}
