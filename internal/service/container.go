package service

import (
	"fiber-starter/internal/repository"
	postService "fiber-starter/internal/service/post"
)

type Container struct {
	PostService *postService.Service
}

func NewServiceContainer(repos *repository.Container) *Container {
	return &Container{
		PostService: postService.NewService(repos),
	}
}
