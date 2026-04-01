package service

import "user_system/internal/repository"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserRepository(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(name, email string) (*repository.User, error) {
	return s.repo.CreateUser(name, email)
}

func (s *UserService) GetUser(id int32) (*repository.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) ListUser() ([]*repository.User, error) {
	return s.repo.ListUser()
}
