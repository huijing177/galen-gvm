package main

import (
	"fmt"
	"galen-gvm/global"
	"galen-gvm/inits"

	_ "go.uber.org/automaxprocs"
)

func main() {
	mainInit()
	fmt.Println("begin")
}

func mainInit() {
	// 读取配置
	// 日志
	// DB
	global.GVA_DB = inits.Gorm()
}
