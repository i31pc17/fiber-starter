package post

import (
	"fiber-starter/internal/service"
	"github.com/gofiber/fiber/v2"
)

type route struct {
	service *service.Container
}

func InitHandler(router fiber.Router, s *service.Container) {
	r := &route{s}

	router.Post("/", r.createPost)
	router.Get("/", r.getAllPosts)
	router.Get("/:id", r.getPostDetail)
	router.Put("/:id", r.updatePostDetail)
	router.Delete("/:id", r.deletePostDetail)
}
