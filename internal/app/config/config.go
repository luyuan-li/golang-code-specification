package config

import (
	"bytes"

	"github.com/spf13/viper"
)

type Config struct {
	App   App   `mapstructure:"app"`
	Mysql Mysql `mapstructure:"mysql"`
}

type (
	App struct {
		Address  string `mapstructure:"address"`
		LogLevel string `mapstructure:"log_level"`
	}

	Mysql struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		Charset  string `mapstructure:"charset"`
		Timezone string `mapstructure:"timezone"`
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
