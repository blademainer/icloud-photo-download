package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	IsChina     bool   `json:"is_china" yaml:"isChina"`
	User        *User  `json:"user" yaml:"user"`
	Debug       bool   `json:"debug" yaml:"debug"`
	LoggerLevel string `json:"logger_level" yaml:"loggerLevel"`
}

type User struct {
	UserName string `json:"user_name" yaml:"userName"`
	Password string `json:"password" yaml:"password"`
}

func ReadConfig(file string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(file)
	// v.SetConfigName("conf")
	v.SetConfigType("yaml")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("icloud")
	v.AutomaticEnv()
	// v.MustBindEnv("user.user_name")
	// v.MustBindEnv("user.password")
	// v.MustBindEnv("loggerlevel")
	// v.Debug()

	// env
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = v.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func confOpt(c *mapstructure.DecoderConfig) {
	c.TagName = "yaml"
}
