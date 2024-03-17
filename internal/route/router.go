package router

import (
	"user-service/internal/module/user/handler"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App, handler *handler.UserHandler) *fiber.App {

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	Api := app.Group("/api")

	// public routes
	v1 := Api.Group("/v1")
	v1.Post("/register", handler.Register)
	v1.Post("/login", handler.Login)
	v1.Get("/user", handler.GetUser)
	v1.Put("/user", handler.UpdateUser)
	v1.Post("/profile", handler.CreateProfile)
	v1.Get("/profile", handler.GetProfile)
	v1.Put("/profile", handler.UpdateProfile)

	private := Api.Group("/private")
	private.Get("/validateToken", handler.GetUser)

	return app

}
