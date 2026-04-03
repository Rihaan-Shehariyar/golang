package service

import (
	"fmt"
	"user_system/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserRepository(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(name, email, password string) (*repository.User, error) {
	return s.repo.CreateUser(name, email, password)
}

func (s *UserService) Login(email, password string) (*repository.User, error) {

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("Invalid Credentials")
	}

	return user, nil

}

func (s *UserService) GetUser(id int32) (*repository.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) ListUser() ([]*repository.User, error) {
	return s.repo.ListUser()
}
