package controller

import (
	"fiber-starter/internal/controller/post"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	post.InitHandler(app.Group("/post"))
}
