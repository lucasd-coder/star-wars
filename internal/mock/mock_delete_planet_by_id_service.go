package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockDeletePlanetByIdService struct {
	mock.Mock
}

func NewMockDeletePlanetByIdService() *MockDeletePlanetByIdService {
	return &MockDeletePlanetByIdService{}
}

func (mock *MockDeletePlanetByIdService) Execute(id string) error {
	args := mock.Called(id)

	var r0 error
	if rf, ok := args.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = args.Error(0)
	}

	return r0
}
