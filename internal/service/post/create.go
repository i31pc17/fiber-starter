package post

import "fiber-starter/internal/entity"

func (p *Service) CreatePost(detail *entity.PostDetail) string {
	p.repos.PostRepo.Create(detail)
	return "asdfasdf"
}
