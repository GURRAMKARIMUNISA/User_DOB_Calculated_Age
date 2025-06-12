package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/GurramKarimunisa/go-user-api/internal/handler" // Adjust import path
	"github.com/GurramKarimunisa/go-user-api/internal/middleware" // Adjust import path
)

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/users")

	// Apply middleware to the API group
	api.Use(middleware.RequestID())
	api.Use(middleware.RequestLogger())

	api.Post("/", userHandler.CreateUser)
	api.Get("/:id", userHandler.GetUserByID)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
	api.Get("/", userHandler.ListUsers)
}