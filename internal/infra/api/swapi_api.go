package api

import (
	"errors"
	"net/http"

	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/infra/api/pool"
	"github.com/lucasd-coder/star-wars/internal/infra/http_client"
	"github.com/lucasd-coder/star-wars/internal/models/external"
	"github.com/lucasd-coder/star-wars/pkg/logger"
)

const (
	uriSwapi = "/api/planets"
)

type SwapiAPI struct {
	HttpConfig *http.Transport
}

func NewSwapiAPI() *SwapiAPI {
	return &SwapiAPI{
		HttpConfig: pool.SwapiAPIConfig(),
	}
}

func (api *SwapiAPI) FindAll() (*external.ResultsSwaApi, error) {
	cfg := config.GetConfig()

	headers := map[string]string{}
	resp, err := http_client.Get(cfg.SwapiUrl+uriSwapi, headers, api.HttpConfig)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		logger.Log.Errorf("swapi integration failure replied with httpStatus: %d", resp.StatusCode)

		return nil, errors.New("unable to parse response from swapi")
	}

	var response external.ResultsSwaApi
	if err := http_client.ParseFromHttpResponse(resp, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
