package main

import (
	"fmt"
	"log"

	"github.com/brendisurfs/brendancreates-backend/email"
	"github.com/brendisurfs/brendancreates-backend/parser"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func submitResponse(c *fiber.Ctx) error {

	var parsedMsg parser.FormSubmit

	fmt.Println("post form data")
	err := c.BodyParser(&parsedMsg)
	if err != nil {
		log.Fatal(err)
	}
	// use mailgun to send the message
	email.SendEmail(parsedMsg.Email, parsedMsg.Subject, parsedMsg.Message)
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
