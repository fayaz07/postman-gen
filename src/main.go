package main

import (
	core "github.com/fayaz07/post-urls-gen/src/core"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	core.InitFlags()
}
