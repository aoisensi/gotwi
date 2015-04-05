package main

import (
	"flag"

	"github.com/google/gxui/drivers/gl"
)

var (
	font = flag.String("font", "", "path to ttf font")
	conf = flag.String("conf", "config.toml", "path to toml config")
)

func main() {
	flag.Parse()
	loadConfig()
	gl.StartDriver(appMain)
}
