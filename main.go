package main

import (
	"fiber-starter/internal/controller"
	"fiber-starter/package/httpserver"
	"fmt"
)

func main() {
	server := httpserver.NewServer()

	controller.InitRouter(server.App)

	if e := server.Listen(); e != nil {
		fmt.Println(e)
	}
}
