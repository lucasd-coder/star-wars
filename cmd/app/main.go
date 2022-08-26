package main

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/app"
	"github.com/lucasd-coder/star-wars/pkg/logger"
)

var cfg config.Config

func main() {
	profile := os.Getenv("GO_PROFILE")
	var path string

	switch profile {
	case "dev":
		path = "./config/config-dev.yml"
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		logger.Log.Fatalf("Config error: %v", err)
	}
	config.ExportConfig(&cfg)

	app.Run(&cfg)
}
