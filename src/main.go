package main

import (
	"github.com/fayaz07/postman-generator/src/core"
)

func main() {
	core.LoadDotEnvConfig()
	core.InitFlags()
	core.ParseFlags()
}
