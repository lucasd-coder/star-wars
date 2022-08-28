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
	plantes, err := service.PlanetRepository.FindAll()
	if err != nil {
		return &[]models.PlanetResponse{}, err
	}

	resultSwapi := service.SwapiIntegrationService.FindAll()

	plantesResponse := make([]models.PlanetResponse, 0, len(plantes))

	for _, plante := range plantes {
		movieAppearances := service.SwapiIntegrationService.SearchMovieAppearances(resultSwapi, &plante)

		plantesResponse = append(plantesResponse, *models.NewPlanetResponse(plante, movieAppearances))
	}

	return &plantesResponse, nil
}
