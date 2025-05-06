package user_http

import (
	"gorm-in-go/pkg/app/user/repository"
	"gorm-in-go/pkg/app/user/service"
	"gorm-in-go/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")
	userHandler := NewUserHandler(service.NewUserService(repository.NewUserRepository(database.GetDB())))

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUserById)
	user.Put("/", userHandler.UpdateUser)
	user.Delete("/", userHandler.DeleteUser)
}
