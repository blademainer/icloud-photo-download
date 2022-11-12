package config

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	_ = os.Setenv("ICLOUD_USER_USERNAME", "test")
	_ = os.Setenv("ICLOUD_USER_PASSWORD", "pass")
	_ = os.Setenv("ICLOUD_LOGGERLEVEL", "debug")
	_ = os.Setenv("ICLOUD_ISCHINA", "false")
	c, err := ReadConfig("testdata/conf.yaml")
	if err != nil {
		return
	}
	js, _ := json.Marshal(c)
	fmt.Println(string(js))
}
