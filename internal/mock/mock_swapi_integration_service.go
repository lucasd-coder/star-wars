package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/models/external"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/stretchr/testify/mock"
)

type MockSwapiIntegrationService struct {
	mock.Mock
	service.CacheService
}

func (mock *MockSwapiIntegrationService) FindAll() *external.ResultsSwaApi {
	args := mock.Called()
	result := args.Get(0)
	return result.(*external.ResultsSwaApi)
}

func (mock *MockSwapiIntegrationService) SearchMovieAppearances(result *external.ResultsSwaApi, planet *models.Planet) int {
	args := mock.Called(result, planet)

	var r0 int
	if rf, ok := args.Get(0).(func(*external.ResultsSwaApi, *models.Planet) int); ok {
		r0 = rf(result, planet)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(int)
		}
	}
	return r0
}
