package main

import (
	"matchtify/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) *fiber.App {
	app.Get("/", handler.GetUsersHandler)
	app.Post("/login", handler.LoginHandler)
	app.Post("/registration", handler.RegistrationHandler)
	app.Post("/forgot-password", handler.ForgotPasswordHandler)

	// POST
	app.Get("/post/create", handler.CreatePostHandler)
	app.Get("/post/delete", handler.DeletePostHandler)
	app.Get("/post/update", handler.UpdatePostHandler)

	// PROFILE
	app.Get("/profile", handler.GetUsersHandler)
	app.Get("/profile/update", handler.GetUsersHandler)

	return app
}
