package conf

import "github.com/dollarkillerx/common/pkg/conf"

var cf *config

func InitConfig(configName string, configPath string) error {
	var c config
	err := conf.InitConfiguration(configName, []string{configPath}, &c)
	cf = &c
	return err
}

func GetConfig() *config {
	if cf == nil {
		panic("Config Uninitialized")
	}

	return cf
}

type config struct {
	ListenAddress string

	PostgresConfiguration *conf.PostgresConfiguration
	RedisConfiguration    *conf.RedisConfiguration
	LoggerConfig          *conf.LoggerConfig
}
