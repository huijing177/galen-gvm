package inits

import (
	"galen-gvm/global"
	"galen-gvm/model/system"

	"go.uber.org/zap"
)

func LoadAllJwtBlackList2Cache() {
	arr, err := system.LoadAllJwtBlackList()
	if err != nil {
		global.GVA_LOG.Error("LoadAllJwtBlackList2Cache err:", zap.Error(err))
	}
	// golang  1.22 for range时，每次迭代v都是新的，不会重复使用之前的地址
	for _, v := range arr {
		global.BlackCache.SetDefault(v, struct{}{})
	}
}
