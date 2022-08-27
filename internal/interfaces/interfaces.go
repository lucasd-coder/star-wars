package interfaces

import (
	"github.com/lucasd-coder/star-wars/internal/models"
)

type (
	PlanetRepository interface {
		Save(planet *models.Planet) error
		FindByName(name string) (*models.Planet, error)
		FindById(id string) (*models.Planet, error)
		FindAll() ([]*models.Planet, error)
		Delete(id string) error
	}

	CreatePlanetService interface {
		Execute(planet *models.PlanetDTO) error
	}

	FindPlanetByIdService interface {
		Execute(id string) (*models.Planet, error)
	}
)
