package service

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lucasd-coder/star-wars/internal/infra/api"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/models/external"
	"github.com/lucasd-coder/star-wars/pkg/logger"
)

type SwapiIntegrationService struct {
	Swapi        interfaces.SwapiAPI
	CacheService interfaces.CacheService
}

func NewSwapiIntegrationService(swapi *api.SwapiAPI, cacheService *CacheService) *SwapiIntegrationService {
	return &SwapiIntegrationService{
		Swapi:        swapi,
		CacheService: cacheService,
	}
}

func (service *SwapiIntegrationService) FindAll() *external.ResultsSwaApi {
	results := &external.ResultsSwaApi{}
	resultCache, err := service.CacheService.Get(context.TODO(), "results")

	if err == redis.Nil {
		logger.Log.Warn("cache with key: results not exists")
	}
	err = json.Unmarshal([]byte(resultCache), &results)

	if err != nil {
		logger.Log.Error(err)
	}

	if len(results.Results) == 0 {
		resp, err := service.Swapi.FindAll()
		if err != nil {
			logger.Log.Error(err.Error())
		}

		results = resp

		err = service.CacheService.Save(context.TODO(), "results", results, 1*time.Hour)
		if err != nil {
			logger.Log.Error(err)
		}

		return results
	}

	return results
}

func (service *SwapiIntegrationService) SearchMovieAppearances(result *external.ResultsSwaApi, planet *models.Planet) int {
	for _, result := range result.Results {
		if strings.Compare(result.Name, planet.Name) == 0 {
			return len(result.Films)
		}
	}
	return 0
}
