//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/pkg/mongodb"
	"github.com/lucasd-coder/star-wars/internal/service"
)

func InitializePlanetRepository() *repository.PlanetRepository {
	wire.Build(config.GetConfig, mongodb.GetClientMongoDB, repository.NewPlanetRepository)
	return &repository.PlanetRepository{}
}

func InitializeCreatePlanetService() *service.CreatePlanetService {
	wire.Build(InitializePlanetRepository, service.NewCreatePlanetService)
	return &service.CreatePlanetService{}
}
