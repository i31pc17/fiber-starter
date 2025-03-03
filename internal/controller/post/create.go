package post

import (
	"github.com/gofiber/fiber/v2"
)

func createPost(c *fiber.Ctx) error {
	req := &createPostRequest{}

	if e := c.QueryParser(req); e != nil {
		return e
	}

	return c.JSON(req)
}
