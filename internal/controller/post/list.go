package post

import "github.com/gofiber/fiber/v2"

func (r *route) getAllPosts(c *fiber.Ctx) error {
	req := &getAllPostsRequest{}

	if err := c.QueryParser(req); err != nil {
		return err
	}

	return c.JSON(req)
}
