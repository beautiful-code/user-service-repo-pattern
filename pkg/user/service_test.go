package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"user-service-repo-pattern/pkg/user/models"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	user := &User{Email: "test@example.com", Password: "password123"}
	mockRepo.On("CreateUser", mock.Anything).Return(user, nil)

	service := NewService(mockRepo)
	result, err := service.CreateUser("test@example.com", "password123")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test@example.com", result.Email)
	mockRepo.AssertExpectations(t)
}
