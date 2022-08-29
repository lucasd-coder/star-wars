package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockFindPlanetByNameService struct {
	mock.Mock
}

func NewMockFindPlanetByNameService() *MockFindPlanetByNameService {
	return &MockFindPlanetByNameService{}
}

func (mock *MockFindPlanetByNameService) Execute(name string) (*models.PlanetResponse, error) {
	args := mock.Called(name)

	r0 := &models.PlanetResponse{}

	if rf, ok := args.Get(0).(func(string) *models.PlanetResponse); ok {
		r0 = rf(name)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*models.PlanetResponse)
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
