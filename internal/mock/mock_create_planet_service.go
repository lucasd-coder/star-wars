package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockCreatePlanetService struct {
	mock.Mock
}

func NewMockCreatePlanetService() *MockCreatePlanetService {
	return &MockCreatePlanetService{}
}

func (mock *MockCreatePlanetService) Execute(planet *models.PlanetDTO) error {
	args := mock.Called(planet)

	var r0 error
	if rf, ok := args.Get(0).(func(*models.PlanetDTO) error); ok {
		r0 = rf(planet)
	} else {
		r0 = args.Error(0)
	}

	return r0
}
