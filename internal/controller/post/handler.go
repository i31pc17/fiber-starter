package post

import "github.com/gofiber/fiber/v2"

func InitHandler(router fiber.Router) {
	router.Post("/", createPost)
	router.Get("/", getAllPosts)
	router.Get("/:id", getPostDetail)
	router.Put("/:id", updatePostDetail)
	router.Delete("/:id", deletePostDetail)
}
