package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/models/external"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FactoryPlanetDTO() *models.PlanetDTO {
	return &models.PlanetDTO{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}
}

func FactoryPlanet() *models.Planet {
	id, _ := primitive.ObjectIDFromHex("630b7bcb419f837457644cbc")
	return &models.Planet{
		ID:      id,
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}
}

func FactoryPlanetResponse() *models.PlanetResponse {
	return &models.PlanetResponse{
		Planet:           *FactoryPlanet(),
		MovieAppearances: 5,
	}
}

func FactoryPlanetsResponse() *[]models.PlanetResponse {
	return &[]models.PlanetResponse{
		{
			Planet:           *FactoryPlanet(),
			MovieAppearances: 5,
		},
	}
}

func FactoryResultsSwaApi() *external.ResultsSwaApi {
	return &external.ResultsSwaApi{
		Results: FactorySwapiPlanet(),
	}
}

func FactorySwapiPlanet() []*external.SwapiPlanet {
	swapiPlanets := make([]*external.SwapiPlanet, 0, 3)

	value1 := &external.SwapiPlanet{
		Name: "Tatooine",
		Films: []string{
			"https://swapi.dev/api/films/1/",
			"https://swapi.dev/api/films/3/",
			"https://swapi.dev/api/films/4/",
			"https://swapi.dev/api/films/5/",
			"https://swapi.dev/api/films/6/",
		},
	}

	value2 := &external.SwapiPlanet{
		Name: "Hoth",
		Films: []string{
			"https://swapi.dev/api/films/2/",
		},
	}

	value3 := &external.SwapiPlanet{
		Name: "Yavin IV",
		Films: []string{
			"https://swapi.dev/api/films/1/",
		},
	}

	swapiPlanets = append(swapiPlanets, value1)
	swapiPlanets = append(swapiPlanets, value2)
	swapiPlanets = append(swapiPlanets, value3)

	return swapiPlanets
}

func FactoryPlanets() []models.Planet {
	value_1, _ := primitive.ObjectIDFromHex("630b7bcb419f837457644cbc")
	value_2, _ := primitive.ObjectIDFromHex("830b7bcb419f837457644cb5")
	value_3, _ := primitive.ObjectIDFromHex("230b7bcb419f837457644cb7")

	return []models.Planet{
		{
			ID:      value_1,
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
		},
		{
			ID:      value_2,
			Name:    "Yavin IV",
			Climate: "temperate, tropical",
			Terrain: "jungle, rainforests",
		},
		{
			ID:      value_3,
			Name:    "Hoth",
			Climate: "frozen",
			Terrain: "tundra, ice caves, mountain ranges",
		},
	}
}
