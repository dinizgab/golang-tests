package mocks

import (
	"github.com/dinizgab/golang-tests/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
    mock.Mock
}

func (m *UserRepositoryMock) FindAll() ([]models.User, error) {
	ret := m.Called()

	var r0 []models.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]models.User)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
