package post

import "fiber-starter/internal/entity"

func (r *Repository) Create(detail *entity.PostDetail) {
	println("create post", detail)
}
