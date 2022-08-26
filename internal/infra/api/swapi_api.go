package api

import (
	"net/http"

	"github.com/lucasd-coder/star-wars/internal/infra/api/pool"
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

func (api *SwapiAPI) FindAll() {}
