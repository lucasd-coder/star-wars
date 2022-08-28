package service_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
	mockTestify "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFailPlanetAlreadyExist(t *testing.T) {
	mockPlanetRepository := new(mock.MockPlanetRepository)

	dto := mock.FactoryPlanetDTO()

	planet := mock.FactoryPlanet()

	mockPlanetRepository.On("FindByName", dto.Name).Return(planet, nil)

	testService := service.CreatePlanetService{mockPlanetRepository}

	err := testService.Execute(dto)

	assert.NotNil(t, err)
	assert.Equal(t, "Planet already exist!", err.Error())
}

func TestCreatePlanetSuccessfully(t *testing.T) {
	mockPlanetRepository := new(mock.MockPlanetRepository)

	dto := mock.FactoryPlanetDTO()

	planet := models.Planet{}

	mockPlanetRepository.On("FindByName", dto.Name).Return(&planet, mongo.ErrNoDocuments)
	mockPlanetRepository.On("Save", mockTestify.Anything).Return(nil)

	testService := service.CreatePlanetService{mockPlanetRepository}

	err := testService.Execute(dto)

	assert.Nil(t, err)
}
