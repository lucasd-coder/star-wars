package service_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
	mockTestify "github.com/stretchr/testify/mock"
)

func TestFindAllCacheableSuccessfully(t *testing.T) {
	swapiAPI := new(mock.MockSwapiAPI)
	cacheService := new(mock.MockCacheService)
	var json string = "{\"results\":[{\"name\":\"Tatooine\",\"films\":[\"https://swapi.dev/api/films/1/\",\"https://swapi.dev/api/films/3/\",\"https://swapi.dev/api/films/4/\",\"https://swapi.dev/api/films/5/\",\"https://swapi.dev/api/films/6/\"]},{\"name\":\"Alderaan\",\"films\":[\"https://swapi.dev/api/films/1/\",\"https://swapi.dev/api/films/6/\"]},{\"name\":\"YavinIV\",\"films\":[\"https://swapi.dev/api/films/1/\"]},{\"name\":\"Hoth\",\"films\":[\"https://swapi.dev/api/films/2/\"]}]}"

	cacheService.On("Get", mockTestify.Anything, mockTestify.Anything).Return(json, nil)

	testService := service.SwapiIntegrationService{swapiAPI, cacheService}

	result := testService.FindAll()

	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func TestFindAllNotCacheableSuccessfully(t *testing.T) {
	swapiAPI := new(mock.MockSwapiAPI)
	cacheService := new(mock.MockCacheService)

	results := mock.FactoryResultsSwaApi()

	cacheService.On("Get", mockTestify.Anything, mockTestify.Anything).Return("", nil)
	cacheService.On("Save", mockTestify.Anything, mockTestify.Anything, mockTestify.Anything, mockTestify.Anything).Return(nil)
	swapiAPI.On("FindAll").Return(results, nil)

	testService := service.SwapiIntegrationService{swapiAPI, cacheService}

	result := testService.FindAll()

	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func TestCountMovieAppearancesSuccessfully(t *testing.T) {
	swapiAPI := new(mock.MockSwapiAPI)
	cacheService := new(mock.MockCacheService)

	results := mock.FactoryResultsSwaApi()

	planet := mock.FactoryPlanet()

	testService := service.SwapiIntegrationService{swapiAPI, cacheService}

	result := testService.SearchMovieAppearances(results, planet)

	assert.Equal(t, result, 5)
}
