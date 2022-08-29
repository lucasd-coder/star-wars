package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockFindPlanetByIdService struct {
	mock.Mock
}

func NewMockFindPlanetByIdService() *MockFindPlanetByIdService {
	return &MockFindPlanetByIdService{}
}

func (mock *MockFindPlanetByIdService) Execute(id string) (*models.PlanetResponse, error) {
	args := mock.Called(id)

	r0 := &models.PlanetResponse{}

	if rf, ok := args.Get(0).(func(string) *models.PlanetResponse); ok {
		r0 = rf(id)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*models.PlanetResponse)
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
