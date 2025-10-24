package apiserver

import (
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// http://localhost:8080/login
	SetupLoginRoute(app)

	// http://localhost:8080/user
	SetupUserRoutes(app)

	// http://localhost:8080/mobile-number
	SetupMobileNumberRoutes(app)
}

func SetupUserRoutes(app *fiber.App) {
	// CRUD User Endpoints
	user := app.Group("/user")
	user.Get("/", handler.GetUserList)
	user.Get("/:id", handler.GetUserByID)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUserByID)
	user.Delete("/:id", handler.DeleteUserByID)
}

func SetupMobileNumberRoutes(app *fiber.App) {
	// Mobile Number Endpoints
	mobile := app.Group("/mobile-number")
	mobile.Get("/", handler.AddMobileNumber)
	mobile.Get("/:number", handler.DeleteMobileNumber)
}

func SetupLoginRoute(app *fiber.App) {
	app.Post("/login", handler.Login)
}
