package main

import (
	"fmt"
	"log"

	"github.com/brendisurfs/brendancreates-backend/email"
	"github.com/brendisurfs/brendancreates-backend/parser"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

}

func submitResponse(c *fiber.Ctx) error {
	fmt.Println("post form data")

	var parsedMsg parser.FormSubmit

	err := c.BodyParser(&parsedMsg)
	if err != nil {
		log.Fatal("could not parse msg: -> ", err)
	}

	err = c.JSON(&parsedMsg)
	if err != nil {
		log.Fatal("could not parse form data as json: -> ", err)
	}

	// use mailgun to send the message
	id, err := email.SendEmail(parsedMsg.Email, parsedMsg.Subject, parsedMsg.Message)
	if err != nil {
		c.SendString("email could not be sent")
	}

	log.Println("success: your id is -> ", id)
	return c.SendString(parsedMsg.Email)
}

func homeHandler(c *fiber.Ctx) error {
	name := "brendi"
	return c.SendString(name)
}

func routesSetup(app *fiber.App) {
	app.Get("/", homeHandler)

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
