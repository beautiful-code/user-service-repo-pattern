package user

import (
	"fmt"
	"regexp"
	"user-service-repo-pattern/pkg/user/models"
	"user-service-repo-pattern/pkg/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(email, password string) (*models.User, error) {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if regex.MatchString(email) {
		return nil, fmt.Errorf("Invalid email format")
	}
	if len(password) <= 6 {
		return nil, fmt.Errorf("Invalid password")
	}
	user := models.User{Email: email, Password: password}
	return s.repo.CreateUser(&user)
}
