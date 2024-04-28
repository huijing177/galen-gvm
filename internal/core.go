package internal

import (
	"fmt"
	"galen-gvm/api"
	"galen-gvm/global"
)

func Run() {
	router := api.Router()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	global.GVA_LOG.Error(router.Run(address).Error())
}
