package main

import (
	"github.com/ErikMora/quest2/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.UrlRoutes(app)
	app.Listen("localhost:8080")
}
