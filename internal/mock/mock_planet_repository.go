package mock

import (
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockPlanetRepository struct {
	mock.Mock
	config.Config
}

func (mock *MockPlanetRepository) Save(planet *models.Planet) error {
	args := mock.Called(planet)

	var r0 error
	if rf, ok := args.Get(0).(func(*models.Planet) error); ok {
		r0 = rf(planet)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (mock *MockPlanetRepository) FindByName(name string) (*models.Planet, error) {
	args := mock.Called(name)

	var r0 *models.Planet
	if rf, ok := args.Get(0).(func(string) *models.Planet); ok {
		r0 = rf(name)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*models.Planet)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (mock *MockPlanetRepository) FindById(id string) (*models.Planet, error) {
	args := mock.Called(id)

	var r0 *models.Planet
	if rf, ok := args.Get(0).(func(string) *models.Planet); ok {
		r0 = rf(id)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*models.Planet)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (mock *MockPlanetRepository) FindAll() ([]models.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]models.Planet), args.Error(0)
}

func (mock *MockPlanetRepository) Delete(id string) (*mongo.DeleteResult, error) {
	args := mock.Called(id)

	var r0 *mongo.DeleteResult
	if rf, ok := args.Get(0).(func(string) *mongo.DeleteResult); ok {
		r0 = rf(id)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*mongo.DeleteResult)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
