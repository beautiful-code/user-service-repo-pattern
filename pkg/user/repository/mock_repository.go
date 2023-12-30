package user

import (
	"user-service-repo-pattern/pkg/user/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) CreateUser(user *models.User) (*models.User, error) {
	args := mock.Called(user)
	return args.Get(0).(*User), args.Error(1)
}
