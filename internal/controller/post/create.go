package post

import (
	"fiber-starter/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (r *route) createPost(c *fiber.Ctx) error {
	req := &createPostRequest{}
	if e := c.QueryParser(req); e != nil {
		return e
	}

	r.service.PostService.CreatePost(&entity.PostDetail{
		Title: req.Title,
		Body:  req.Body,
	})

	return c.JSON(req)
}
