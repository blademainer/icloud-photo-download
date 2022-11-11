package config

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	IsChina bool  `json:"is_china" yaml:"isChina"`
	User    *User `json:"user" yaml:"user"`
	Debug   bool  `json:"debug" yaml:"debug"`
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
	// v.SetEnvPrefix("icloud")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// v.AutomaticEnv()
	// viper.MustBindEnv("user.user_name", "ICLOUD_USER_NAME")
	// viper.MustBindEnv("user.password", "ICLOUD_PASSWORD")
	v.Debug()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	js, _ := json.Marshal(c)
	fmt.Println(string(js))
	err = v.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func confOpt(c *mapstructure.DecoderConfig) {
	c.TagName = "yaml"
}

// defaultDecoderConfig returns default mapsstructure.DecoderConfig with suppot
// of time.Duration values & string slices
func defaultDecoderConfig(
	output interface{}, opts ...viper.DecoderConfigOption,
) *mapstructure.DecoderConfig {
	c := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
