package post

import "fiber-starter/internal/repository"

type Service struct {
	repos *repository.Container
}

func NewService(repos *repository.Container) *Service {
	return &Service{repos: repos}
}
