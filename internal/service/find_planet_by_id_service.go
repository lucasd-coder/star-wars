package service

import (
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
)

type FindPlanetByIdService struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewFindPlanetByIdService(planetRepository *repository.PlanetRepository) *FindPlanetByIdService {
	return &FindPlanetByIdService{
		PlanetRepository: planetRepository,
	}
}

func (service *FindPlanetByIdService) Execute(id string) (*models.Planet, error) {
	planet, err := service.PlanetRepository.FindById(id)
	if err != nil {
		return &models.Planet{}, err
	}

	return planet, nil
}
