package service

import (
	"github.com/lucasd-coder/star-wars/internal/errs"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeletePlanetByIdService struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewDeletePlanetByIdService(planetRepository *repository.PlanetRepository) *DeletePlanetByIdService {
	return &DeletePlanetByIdService{
		PlanetRepository: planetRepository,
	}
}

func (service *DeletePlanetByIdService) Execute(id string) error {
	if !primitive.IsValidObjectID(id) {
		return &errs.AppError{
			Code:    400,
			Message: "id in not valid format",
		}
	}

	result, err := service.PlanetRepository.Delete(id)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return &errs.AppError{
			Code:    404,
			Message: "planet not found",
		}
	}

	return nil
}
