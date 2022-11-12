package main

import (
	"encoding/json"
	"fmt"
	"github.com/blademainer/icloud-photo-download/pkg/icloud"
	"github.com/blademainer/icloud-photo-download/pkg/icloud/config"
)

func main() {
	// _ = os.Setenv("ICLOUD_USER_USERNAME", "your_user")
	// _ = os.Setenv("ICLOUD_USER_PASSWORD", "your_pass")
	c, err := config.ReadConfig("conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	js, _ := json.Marshal(c)
	fmt.Println(string(js))

	cloud, err := icloud.NewICloud(
		c,
	)
	if err != nil {
		panic(err)
		return
	}
	err = cloud.Login()
	if err != nil {
		panic(err)
		return
	}
}
