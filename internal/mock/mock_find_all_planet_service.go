package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockFindAllPlanetService struct {
	mock.Mock
}

func NewMockFindAllPlanetService() *MockFindAllPlanetService {
	return &MockFindAllPlanetService{}
}

func (mock *MockFindAllPlanetService) Execute() (*[]models.PlanetResponse, error) {
	args := mock.Called()

	var r0 *[]models.PlanetResponse
	if rf, ok := args.Get(0).(func() *[]models.PlanetResponse); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*[]models.PlanetResponse)
		}
	}
	var r1 error
	if rf, ok := args.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
