package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func submitResponse(c *fiber.Ctx) error {
	fmt.Println("post form data")
	data := c.Body()
	fmt.Println(string(data))
	return c.SendString("post form data")
}

func routesSetup(app *fiber.App) {
	app.Post("/api/submit", submitResponse)
}

func main() {
	app := fiber.New()

	// middleware
	app.Use(cors.New())

	routesSetup(app)

	// listener
	app.Listen(":8080")
}
