package controller

import (
	"fiber-starter/internal/controller/post"
	"fiber-starter/internal/repository"
	"fiber-starter/internal/service"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	RDB := "RDB"
	Redis := "Redis"
	repos := repository.NewRepositoryContainer(RDB, Redis)
	s := service.NewServiceContainer(repos)

	post.InitHandler(app.Group("/post"), s)
}
