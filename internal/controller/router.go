package controller

import (
	"fiber-starter/internal/controller/post"
	"fiber-starter/internal/service"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	s := service.NewService()

	post.InitHandler(app.Group("/post"), s)
}
