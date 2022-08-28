package interfaces

import (
	"context"
	"time"

	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/models/external"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PlanetRepository interface {
		Save(planet *models.Planet) error
		FindByName(name string) (*models.Planet, error)
		FindById(id string) (*models.Planet, error)
		FindAll() ([]models.Planet, error)
		Delete(id string) (*mongo.DeleteResult, error)
	}

	CreatePlanetService interface {
		Execute(planet *models.PlanetDTO) error
	}

	FindPlanetByIdService interface {
		Execute(id string) (*models.PlanetResponse, error)
	}

	FindPlanetByNameService interface {
		Execute(name string) (*models.PlanetResponse, error)
	}

	SwapiIntegrationService interface {
		FindAll() *external.ResultsSwaApi
		SearchMovieAppearances(result *external.ResultsSwaApi, planet *models.Planet) int
	}

	CacheService interface {
		Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error
		Get(ctx context.Context, key string) (string, error)
	}

	FindAllPlanetService interface {
		Execute() (*[]models.PlanetResponse, error)
	}

	DeletePlanetByIdService interface {
		Execute(id string) error
	}
)
