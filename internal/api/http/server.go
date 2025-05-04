package server

import "github.com/gofiber/fiber/v2"

func StartServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":8080")
}
