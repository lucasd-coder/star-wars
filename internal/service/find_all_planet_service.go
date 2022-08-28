package service

import (
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
)

type FindAllPlanetService struct {
	SwapiIntegrationService interfaces.SwapiIntegrationService
	PlanetRepository        interfaces.PlanetRepository
}

func NewFindAllPlanetService(swapi *SwapiIntegrationService,
	planetRepository *repository.PlanetRepository,
) *FindAllPlanetService {
	return &FindAllPlanetService{
		SwapiIntegrationService: swapi,
		PlanetRepository:        planetRepository,
	}
}

func (service *FindAllPlanetService) Execute() (*[]models.PlanetResponse, error) {
	planets, err := service.PlanetRepository.FindAll()
	if err != nil {
		return &[]models.PlanetResponse{}, err
	}

	resultSwapi := service.SwapiIntegrationService.FindAll()

	planetesResponse := make([]models.PlanetResponse, 0, len(planets))

	for _, planete := range planets {
		movieAppearances := service.SwapiIntegrationService.SearchMovieAppearances(resultSwapi, &planete)

		planetesResponse = append(planetesResponse, *models.NewPlanetResponse(planete, movieAppearances))
	}

	return &planetesResponse, nil
}
