package service

import (
	"BWG/entity"
	"BWG/internal/repo"
	"fmt"
)

type UserService struct {
	repo repo.User
}

func NewUserService(repo repo.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entity.User) (int, error) {
	id, err := s.repo.CreateUser(user)
	return id, err
}

func (s *UserService) GetUser(userId int) (entity.User, error) {
	user, err := s.repo.GetUser(userId)
	return user, err
}

func (s *UserService) GetUsers(options entity.Options) ([]entity.User, error) {
	offset := (options.Page - 1) * options.PerPage
	query := fmt.Sprintf("SELECT * FROM users LIMIT %d OFFSET %d", options.PerPage, offset)
	users, err := s.repo.GetUsers(query)

	return users, err
}
