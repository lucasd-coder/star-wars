package service_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFailPlanetByNomeNotFound(t *testing.T) {
	swapiIntegrationService := new(mock.MockSwapiIntegrationService)
	mockPlanetRepository := new(mock.MockPlanetRepository)

	var name string = "name not found"
	planet := models.Planet{}

	mockPlanetRepository.On("FindByName", name).Return(&planet, mongo.ErrNoDocuments)

	testService := service.FindPlanetByNameService{swapiIntegrationService, mockPlanetRepository}

	_, err := testService.Execute(name)
	assert.NotNil(t, err)
	assert.Equal(t, "planet not found", err.Error())
}

func TestFindPlanetByNameSuccessfully(t *testing.T) {
	swapiIntegrationService := new(mock.MockSwapiIntegrationService)
	mockPlanetRepository := new(mock.MockPlanetRepository)

	planet := mock.FactoryPlanet()
	resultsSwaApi := mock.FactoryResultsSwaApi()

	var name string = "Tatooine"

	mockPlanetRepository.On("FindByName", name).Return(planet, nil)
	swapiIntegrationService.On("FindAll").Return(resultsSwaApi)
	swapiIntegrationService.On("SearchMovieAppearances", resultsSwaApi, planet).Return(3)

	testService := service.FindPlanetByNameService{swapiIntegrationService, mockPlanetRepository}

	resp, err := testService.Execute(name)
	assert.Nil(t, err)
	assert.Equal(t, resp.ID, planet.ID)
	assert.Equal(t, resp.Terrain, planet.Terrain)
}
