package core

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	foo     bool
	Version string
}

var appConfig Config

func LoadDotEnvConfig() {
	godotenv.Load()

	appConfig.Version = os.Getenv("version")

}
