package app

import (
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/pkg/logger"
)

func Run(cfg *config.Config) {
	// Log config
	logger.SetUpLog(cfg)
}
