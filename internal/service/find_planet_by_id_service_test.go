package service_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFailIdInvalid(t *testing.T) {
	swapiIntegrationService := new(mock.MockSwapiIntegrationService)
	mockPlanetRepository := new(mock.MockPlanetRepository)

	var id string = "id invalid"

	testService := service.FindPlanetByIdService{swapiIntegrationService, mockPlanetRepository}

	_, err := testService.Execute(id)
	assert.NotNil(t, err)
	assert.Equal(t, "id in not valid format", err.Error())
}

func TestFailPlanetByIdNotFound(t *testing.T) {
	swapiIntegrationService := new(mock.MockSwapiIntegrationService)
	mockPlanetRepository := new(mock.MockPlanetRepository)

	var id string = "630b7bcb419f837457644cbc"
	plant := models.Planet{}

	mockPlanetRepository.On("FindById", id).Return(&plant, mongo.ErrNoDocuments)

	testService := service.FindPlanetByIdService{swapiIntegrationService, mockPlanetRepository}

	_, err := testService.Execute(id)
	assert.NotNil(t, err)
	assert.Equal(t, "planet not found", err.Error())
}

func TestFindPlanetByIdSuccessfully(t *testing.T) {
	swapiIntegrationService := new(mock.MockSwapiIntegrationService)
	mockPlanetRepository := new(mock.MockPlanetRepository)

	planet := mock.FactoryPlanet()
	resultsSwaApi := mock.FactoryResultsSwaApi()

	var id string = "630b7bcb419f837457644cbc"

	mockPlanetRepository.On("FindById", id).Return(planet, nil)
	swapiIntegrationService.On("FindAll").Return(resultsSwaApi)
	swapiIntegrationService.On("SearchMovieAppearances", resultsSwaApi, planet).Return(3)

	testService := service.FindPlanetByIdService{swapiIntegrationService, mockPlanetRepository}

	resp, err := testService.Execute(id)
	assert.Nil(t, err)
	assert.Equal(t, resp.ID, planet.ID)
	assert.Equal(t, resp.Terrain, planet.Terrain)
}
