package post

import "fiber-starter/internal/entity"

type post struct{}

type Post interface {
	CreatePost(detail *entity.PostDetail) string
}

func NewPost() Post {
	return &post{}
}
