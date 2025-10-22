package apiserver

import (
	"log"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/handler"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Use(middleware.RequestLogger())

	app.Post("/login", handler.Login) 
	app.Post("/user", handler.CreateUser) 

	protected := app.Group("/", middleware.Protected())

	protected.Get("/user", handler.GetUserList)
	protected.Get("/user/:id", handler.GetUserByID)
	protected.Put("/user/:id", handler.UpdateUserByID)
	protected.Delete("/user/:id", handler.DeleteUserByID)

	protected.Post("/user/:id/mobile-number", handler.AddMobileNumber)
	protected.Delete("/user/:id/mobile-number/:number", handler.DeleteMobileNumber)

	log.Println(app.Listen(":8080"))
}
