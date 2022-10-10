package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/cloudflare/cloudflare-go"
	"github.com/go-bai/cloudflare-ddns/app"
)

func main() {
	version := flag.Bool("version", false, "a bool")
	init := flag.Bool("init", false, "a bool")
	run := flag.Bool("run", false, "a bool")
	flag.Parse()

	switch {
	case *version:
		// print version
		fmt.Println("version 0.0.1", runtime.GOOS, runtime.GOARCH)
	case *init:
		// init
		fmt.Println("start init")
		core := &app.Core{}
		core.Init()
	case *run:
		// run
		fmt.Println("start run")
		config := &app.Config{}
		err := config.Decode()
		if err != nil {
			log.Fatal(err)
		}
		api, err := cloudflare.New(config.Cloudflare.CfApiKey, config.Cloudflare.CfApiEmail)
		if err != nil {
			log.Fatalf("cloudflare new error: %v", err)
		}
		core := app.NewCore(config, api)
		core.Run()
	default:
		fmt.Println("use -h")
	}

}
