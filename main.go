package main

import (
	"fmt"

	"github.com/brendisurfs/brendancreates-backend/parser"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func submitResponse(c *fiber.Ctx) error {
	fmt.Println("post form data")
	data := c.Body()

	// convert to json
	formMsg := parser.MessageParser(data)
	fmt.Println(formMsg.Email)

	return c.SendString("form submitted")
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
