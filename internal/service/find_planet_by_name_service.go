package service

import (
	"github.com/lucasd-coder/star-wars/internal/errs"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindPlanetByNameService struct {
	SwapiIntegrationService interfaces.SwapiIntegrationService
	PlanetRepository        interfaces.PlanetRepository
}

func NewFindPlanetByNameService(swapi *SwapiIntegrationService,
	planetRepository *repository.PlanetRepository,
) *FindPlanetByNameService {
	return &FindPlanetByNameService{
		SwapiIntegrationService: swapi,
		PlanetRepository:        planetRepository,
	}
}

func (service *FindPlanetByNameService) Execute(name string) (*models.PlanetResponse, error) {
	planet, err := service.PlanetRepository.FindByName(name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.PlanetResponse{}, &errs.AppError{
				Code:    404,
				Message: "planet not found",
			}
		}
		return &models.PlanetResponse{}, err
	}

	result := service.SwapiIntegrationService.FindAll()

	movieAppearances := service.SwapiIntegrationService.SearchMovieAppearances(result, planet)

	planetResponse := models.NewPlanetResponse(*planet, movieAppearances)

	return planetResponse, nil
}
