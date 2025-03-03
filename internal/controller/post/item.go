package post

import "github.com/gofiber/fiber/v2"

func getPostDetail(c *fiber.Ctx) error {
	req := &getPostDetailRequest{}

	if err := c.ParamsParser(req); err != nil {
		return err
	}

	return c.JSON(req)
}

func updatePostDetail(c *fiber.Ctx) error {
	req := &updatePostDetailRequest{}

	if err := c.ParamsParser(req); err != nil {
		return err
	}

	if err := c.QueryParser(req); err != nil {
		return err
	}

	return c.JSON(req)
}

func deletePostDetail(c *fiber.Ctx) error {
	req := &deletePostDetailRequest{}

	if err := c.ParamsParser(req); err != nil {
		return err
	}

	return c.JSON(req)
}
