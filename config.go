package system

import (
	"go2api/system/config"
	"go2api/system/config/models"
	"os"
)

func NewGo2Config(filePath string) *config.Go2Config {
	var go2conf *config.Go2Config = &config.Go2Config{}
	go2conf.EnvVars = make(map[string]string)
	var config models.Config = config.LoadConfigFile(filePath)
	for key, val := range config.EnvVars {
		go2conf.EnvVars[key] = os.Getenv(val)
	}
	go2conf.Config = config
	return go2conf
}
