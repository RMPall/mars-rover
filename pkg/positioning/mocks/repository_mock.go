package mocks

import (
	"github.com/stretchr/testify/mock"
	"mars-rover/pkg/positioning"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) Save(robot *positioning.Robot) error {
	args := r.Called(robot)
	return args.Error(0)
}
