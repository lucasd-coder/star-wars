package service

import (
	"github.com/lucasd-coder/star-wars/internal/errs"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreatePlanetService struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewCreatePlanetService(planetRepository *repository.PlanetRepository) *CreatePlanetService {
	return &CreatePlanetService{
		PlanetRepository: planetRepository,
	}
}

func (createPlanet *CreatePlanetService) Execute(dto *models.PlanetDTO) error {
	existentPlanet, err := createPlanet.PlanetRepository.FindByName(dto.Name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			planet := models.NewPlanet(*dto)
			err = createPlanet.PlanetRepository.Save(planet)
			if err != nil {
				return err
			}
		}
	}

	if existentPlanet.Name != "" {
		return &errs.AppError{
			Message: "Planet already exist!",
			Code:    400,
		}
	}

	return err
}
