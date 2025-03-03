package service

import "fiber-starter/internal/service/post"

type service struct{}

type Service interface {
	post.Post
}

func NewService() Service {
	return &service{}
}
