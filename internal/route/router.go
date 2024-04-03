package router

import (
	"user-service/internal/module/user/handler"
	"user-service/internal/pkg/helpers/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App, handler *handler.UserHandler, m *middleware.Middleware) *fiber.App {

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	Api := app.Group("/api")

	// public routes
	v1 := Api.Group("/v1")
	v1.Post("/register", handler.Register)
	v1.Post("/login", handler.Login)
	v1.Get("/user", m.VerifyBearerToken, handler.GetUser)
	v1.Put("/user", m.VerifyBearerToken, handler.UpdateUser)
	v1.Post("/profile", m.VerifyBearerToken, handler.CreateProfile)
	v1.Get("/profile", m.VerifyBearerToken, handler.GetProfile)
	v1.Put("/profile", m.VerifyBearerToken, handler.UpdateProfile)

	private := Api.Group("/private")
	private.Get("user/validate", handler.ValidateToken)
	private.Get("user/profile", handler.GetProfilePrivate)

	return app

}
