package main

import (
	"flag"
	"fmt"
	"runtime"

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
		core := app.NewCore()
		core.Run()
	default:
		fmt.Println("use -h")
	}

}
