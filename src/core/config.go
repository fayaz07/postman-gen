package core

import (
	"os"
	"strings"

	"github.com/fayaz07/postman-generator/src/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	initialized bool
	Version     string
	Languages   []string
}

var appConfig Config

func GetAppConfig() *Config {
	if !appConfig.initialized {
		LoadDotEnvConfig()
	}
	return &appConfig
}

func LoadDotEnvConfig() {
	godotenv.Load()

	appConfig.Version = os.Getenv(utils.VERSION_KEY)
	appConfig.initialized = true
	appConfig.Languages = strings.Split(os.Getenv(utils.LANGUAGES_KEY), ",")
}
