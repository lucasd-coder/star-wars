package service

import (
	"github.com/lucasd-coder/star-wars/internal/errs"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindPlanetByIdService struct {
	SwapiIntegrationService interfaces.SwapiIntegrationService
	PlanetRepository        interfaces.PlanetRepository
}

func NewFindPlanetByIdService(swapi *SwapiIntegrationService,
	planetRepository *repository.PlanetRepository,
) *FindPlanetByIdService {
	return &FindPlanetByIdService{
		SwapiIntegrationService: swapi,
		PlanetRepository:        planetRepository,
	}
}

func (service *FindPlanetByIdService) Execute(id string) (*models.PlanetResponse, error) {
	if !primitive.IsValidObjectID(id) {
		return &models.PlanetResponse{}, &errs.AppError{
			Code:    400,
			Message: "id in not valid format",
		}
	}

	planet, err := service.PlanetRepository.FindById(id)
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
