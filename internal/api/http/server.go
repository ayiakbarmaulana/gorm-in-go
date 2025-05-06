package server

import (
	user_http "gorm-in-go/internal/api/http/user"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	user_http.UserRoutes(app)

	app.Listen(":8080")
}
