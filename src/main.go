package main

import (
	"github.com/fayaz07/postman-gen/src/core"
)

func main() {
	core.LoadDotEnvConfig()
	core.InitFlags()
	core.ParseFlags()
}
