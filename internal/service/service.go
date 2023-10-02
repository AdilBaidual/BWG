package service

import (
	"BWG/internal/repo"
)

type Service struct {
}

func NewService(repo *repo.Repository) *Service {
	return &Service{}
}
