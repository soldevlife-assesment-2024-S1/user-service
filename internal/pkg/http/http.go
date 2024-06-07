package http

import (
	"log"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupHttpEngine() *fiber.App {
	// init http server
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}),
	)

	// setup tracer
	app.Use(otelfiber.Middleware())

	// setup logger
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${TagReqHeaders} ${body} ${error} \n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	return app

}

func StartHttpServer(app *fiber.App, port string) {
	log.Fatal(app.Listen(":" + port))
}
