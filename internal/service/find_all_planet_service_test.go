package service_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
	mockTestify "github.com/stretchr/testify/mock"
)

func TestFindAllPlanetSuccessfully(t *testing.T) {
	swapiIntegrationService := new(mock.MockSwapiIntegrationService)
	mockPlanetRepository := new(mock.MockPlanetRepository)

	planet := mock.FactoryPlanets()
	resultsSwaApi := mock.FactoryResultsSwaApi()

	mockPlanetRepository.On("FindAll").Return(planet, nil)
	swapiIntegrationService.On("FindAll").Return(resultsSwaApi)
	swapiIntegrationService.On("SearchMovieAppearances", mockTestify.Anything, mockTestify.Anything).Return(3)

	testService := service.FindAllPlanetService{swapiIntegrationService, mockPlanetRepository}

	resp, err := testService.Execute()
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}
