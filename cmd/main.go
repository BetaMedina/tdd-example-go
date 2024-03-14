package main

import (
	"tdd/api/router"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	router.HealthCheckRouter(app)
	router.StatementRouter(app)
	app.Listen(":8000")
}
