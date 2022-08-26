package pool

import (
	"net/http"
	"time"

	"github.com/lucasd-coder/star-wars/config"
)

func SwapiAPIConfig() *http.Transport {
	cfg := config.GetConfig()

	connTimeout, err := time.ParseDuration(cfg.SwapiApiConnTimeout)
	if err != nil {
		connTimeout = 60 * time.Second
	}

	readTimeout, err := time.ParseDuration(cfg.SwapiApiReadTimeout)
	if err != nil {
		readTimeout = 60 * time.Second
	}

	return &http.Transport{
		MaxIdleConns:          cfg.SwapiApiMaxConn,
		IdleConnTimeout:       connTimeout,
		MaxConnsPerHost:       cfg.SwapiApiMaxRoutes,
		ResponseHeaderTimeout: readTimeout,
	}
}
