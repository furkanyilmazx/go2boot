package models

import "reflect"

type Config struct {
	DB          `json:"db"`
	App         `json:"app"`
	Vars        map[string]string `json:"vars"`
	EnvVars     map[string]string `json:"envVars"`
	StaticFiles []StaticFile      `json:"staticFiles"`
}

func (c Config) IsEmpty() bool {
	return reflect.DeepEqual(Config{}, c)
}
