package main

import (
	"flag"
	"log"

	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/dollarkillerx/zim/internal/manager/conf"
	"github.com/dollarkillerx/zim/internal/manager/server"
)

// 管理端

/**
职责：
- 基础api 管理 鉴权
- 用户在线控制
*/

var (
	configName string
	configPath string
)

func init() {
	flag.StringVar(&configName, "configName", "config", "config name")
	flag.StringVar(&configPath, "configPath", "configs", "config path")
	flag.Parse()
}

func main() {
	initializeBasicDependencies()

	managerServer, err := server.NewManagerServer()
	if err != nil {
		logger.Panicf("NewManagerServer Run Error: %s", err)
	}

	if err := managerServer.Run(); err != nil {
		logger.Panicf("ManagerServer Run Error: %s", err)
	}
}

func initializeBasicDependencies() {
	// init conf
	err := conf.InitConfig(configName, configPath)
	if err != nil {
		log.Fatalf("initializeBasicDependencies init config error: %s \n", err)
	}

	// init logger
	logger.InitLogger(*conf.GetConfig().LoggerConfig)
}
