package main

import (
	"github.com/blademainer/icloud-photo-download/pkg/icloud"
	"github.com/blademainer/icloud-photo-download/pkg/icloud/config"
)

func main() {
	c, err := config.ReadConfig("conf/conf.yaml.tmp")
	if err != nil {
		panic(err)
	}
	cloud, err := icloud.NewICloud(
		c,
	)
	if err != nil {
		panic(err)
		return
	}
	cloud.Login()
}
