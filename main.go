package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blademainer/icloud-photo-download/pkg/icloud"
	"github.com/blademainer/icloud-photo-download/pkg/icloud/config"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// shutdown functions
	shutdownFunctions := make([]func(context.Context), 0)

	// signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	// errgroup
	g, ctx := errgroup.WithContext(ctx)

	g.Go(
		func() error {
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
				return err
			}
			err = cloud.Login()
			if err != nil {
				return err
			}
			return err
		},
	)

	// http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	_, err := writer.Write([]byte("hello!\n"))
	//	if err != nil {
	//		fmt.Printf("write error: %v \n", err.Error())
	//	}
	// })
	// err := http.ListenAndServe(":8080", nil)
	select {
	case <-ctx.Done():
		break
	case <-interrupt:
		break
	}

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	for _, shutdown := range shutdownFunctions {
		shutdown(timeout)
	}
	err := g.Wait()
	if err != nil {
		panic(err)
	}

}
