package httpserver

import "github.com/gofiber/fiber/v2"

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	server := &Server{}

	app := fiber.New()
	server.App = app

	return server
}

func (server *Server) Listen() error {
	return server.App.Listen(":3000")
}
