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

    return ret.Get(0).([]models.User), ret.Get(1).(error)
}

func (m *UserRepositoryMock) FindByID(id string) (models.User, error) {
    ret := m.Called(id)

    return ret.Get(0).(models.User), ret.Get(1).(error)
}

func (m *UserRepositoryMock) Save(user models.User) error {
    ret := m.Called(user)

    return ret.Get(0).(error)
}

func (m *UserRepositoryMock) FollowUser(userId string, followUserId string) error {
    ret := m.Called(userId, followUserId)

    return ret.Get(0).(error)
}
