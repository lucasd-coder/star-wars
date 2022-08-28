package mock

import (
	"github.com/lucasd-coder/star-wars/internal/models/external"
	"github.com/stretchr/testify/mock"
)

type MockSwapiAPI struct {
	mock.Mock
}

func (mock *MockSwapiAPI) FindAll() (*external.ResultsSwaApi, error) {
	args := mock.Called()

	var r0 *external.ResultsSwaApi
	if rf, ok := args.Get(0).(func() *external.ResultsSwaApi); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*external.ResultsSwaApi)
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
