//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/infra/api"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/pkg/mongodb"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/lucasd-coder/star-wars/pkg/cache"
)

func InitializePlanetRepository() *repository.PlanetRepository {
	wire.Build(config.GetConfig, mongodb.GetClientMongoDB, repository.NewPlanetRepository)
	return &repository.PlanetRepository{}
}

func InitializeCreatePlanetService() *service.CreatePlanetService {
	wire.Build(InitializePlanetRepository, service.NewCreatePlanetService)
	return &service.CreatePlanetService{}
}

func InitializeSwapiAPI() *api.SwapiAPI {
	wire.Build(api.NewSwapiAPI)
	return &api.SwapiAPI{}
}

func InitializeSwapiIntegrationService() *service.SwapiIntegrationService {
	wire.Build(InitializeSwapiAPI, InitializeCacheService, service.NewSwapiIntegrationService)
	return &service.SwapiIntegrationService{}
}

func InitializeFindPlanetByIdService() *service.FindPlanetByIdService {
	wire.Build(InitializeSwapiIntegrationService, InitializePlanetRepository, service.NewFindPlanetByIdService)
	return &service.FindPlanetByIdService{}
}

func InitializeCacheService() *service.CacheService {
	wire.Build(cache.GetClient, service.NewCacheService)
	return &service.CacheService{}
}

func InitializeFindPlanetByNameService() *service.FindPlanetByNameService {
	wire.Build(InitializeSwapiIntegrationService, InitializePlanetRepository, service.NewFindPlanetByNameService)
	return &service.FindPlanetByNameService{}
}

func InitializeFindAllPlanetService() *service.FindAllPlanetService {
	wire.Build(InitializeSwapiIntegrationService, InitializePlanetRepository, service.NewFindAllPlanetService)
	return &service.FindAllPlanetService{}
}

func InitializeDeletePlanetByIdService() *service.DeletePlanetByIdService {
	wire.Build(InitializePlanetRepository, service.NewDeletePlanetByIdService)
	return &service.DeletePlanetByIdService{}
}
