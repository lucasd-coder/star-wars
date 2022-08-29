package models_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPlanet(t *testing.T) {
	dto := models.PlanetDTO{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}

	planet := models.NewPlanet(dto)

	require.Equal(t, planet.Name, dto.Name)
	require.Equal(t, planet.Climate, dto.Climate)
	require.Equal(t, planet.Terrain, dto.Terrain)
}

func TestModel_NewPlanetResponse(t *testing.T) {
	planet := models.Planet{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}

	test := models.NewPlanetResponse(planet, 5)

	require.Equal(t, planet.Name, test.Name)
	require.Equal(t, planet.Climate, test.Climate)
	require.Equal(t, planet.Terrain, test.Terrain)
	require.Equal(t, 5, test.MovieAppearances)
}
