package config

import (
	"bytes"

	"github.com/spf13/viper"
)

type Config struct {
	App App `mapstructure:"app"`
}

type (
	App struct {
		Addr     string `mapstructure:"address"`
		LogLevel string `mapstructure:"log_level"`
	}
)

func ReadConfig(data []byte) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	reader := bytes.NewReader(data)
	err := v.ReadConfig(reader)
	if err != nil {
		return nil, err
	}
	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
