package inits

import (
	"flag"
	"fmt"
	"os"

	"galen-gvm/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper() *viper.Viper {
	var configPath string

	flag.StringVar(&configPath, "c", "", "choose config file.")
	flag.Parse()
	if configPath == "" { // 判断命令行参数是否为空
		if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				configPath = global.ConfigDefaultFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), global.ConfigDefaultFile)
			case gin.ReleaseMode:
				configPath = global.ConfigReleaseFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), global.ConfigReleaseFile)
			case gin.TestMode:
				configPath = global.ConfigTestFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), global.ConfigTestFile)
			}
		} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
			configPath = configEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", global.ConfigEnv, configPath)
		}
	} else { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configPath)
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}

	return v
}
