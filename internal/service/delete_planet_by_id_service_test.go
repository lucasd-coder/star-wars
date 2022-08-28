package service_test

import (
	"testing"

	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFailDeleteIdInvalid(t *testing.T) {
	mockPlanetRepository := new(mock.MockPlanetRepository)

	var id string = "id invalid"

	testService := service.DeletePlanetByIdService{mockPlanetRepository}

	err := testService.Execute(id)
	assert.NotNil(t, err)
	assert.Equal(t, "id in not valid format", err.Error())
}

func TestFailDeleteByIdNotFound(t *testing.T) {
	mockPlanetRepository := new(mock.MockPlanetRepository)

	var id string = "630b7bcb419f837457644cbc"

	mongoResult := mongo.DeleteResult{
		DeletedCount: 0,
	}

	mockPlanetRepository.On("Delete", id).Return(&mongoResult, nil)

	testService := service.DeletePlanetByIdService{mockPlanetRepository}

	err := testService.Execute(id)
	assert.NotNil(t, err)
	assert.Equal(t, "planet not found", err.Error())
}

func TestFailDeleteSuccessfully(t *testing.T) {
	mockPlanetRepository := new(mock.MockPlanetRepository)

	var id string = "630b7bcb419f837457644cbc"

	mongoResult := mongo.DeleteResult{
		DeletedCount: 1,
	}

	mockPlanetRepository.On("Delete", id).Return(&mongoResult, nil)

	testService := service.DeletePlanetByIdService{mockPlanetRepository}

	err := testService.Execute(id)
	assert.Nil(t, err)
}
