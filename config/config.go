package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig(path, name string) {
	if path != "" {
		viper.AddConfigPath(path)
		viper.SetConfigName(name)
	} else {
		viper.AddConfigPath("./config")
		viper.AddConfigPath("./")
		viper.SetConfigName("conf")
	}

	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config.LoadConfig: failed to read config. path:%s, err: %v", path, err))
	}
}
