package config

import (
	"encoding/json"
	"fmt"
	"go2api/system/config/models"
	"os"
)

type Go2Config struct {
	models.Config
	EnvVars map[string]string
}

func (g Go2Config) GetEnv(key string) string {
	fmt.Println(g.EnvVars[key])
	return g.EnvVars[key]
}

func (g *Go2Config) SetEnv(key, val string) {
	fmt.Println(key)
	err := os.Setenv(key, val)
	if err != nil {
		fmt.Println(err.Error())
	}
	g.EnvVars[key] = os.Getenv(key)
}

func LoadConfigFile(root string) models.Config {
	var config models.Config
	configFile, err := os.Open(root)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}
