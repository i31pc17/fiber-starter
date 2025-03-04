package repository

import "fiber-starter/internal/repository/post"

type Container struct {
	PostRepo *post.Repository
}

func NewRepositoryContainer(RDB interface{}, Redis interface{}) *Container {
	PostRepo := post.NewPostRepository(RDB, Redis)
	return &Container{
		PostRepo,
	}
}
