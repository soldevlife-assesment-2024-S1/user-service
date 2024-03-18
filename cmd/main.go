package main

import (
	"user-service/config"
	"user-service/internal/module/user/handler"
	"user-service/internal/module/user/repositories"
	"user-service/internal/module/user/usecases"
	"user-service/internal/pkg/database"
	"user-service/internal/pkg/helpers/middleware"
	"user-service/internal/pkg/http"
	"user-service/internal/pkg/log"
	router "user-service/internal/route"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.InitConfig()

	app := initService(cfg)

	// start http server
	http.StartHttpServer(app, cfg.HttpServer.Port)
}

func initService(cfg *config.Config) *fiber.App {
	db := database.GetConnection(&cfg.Database)
	// redis := redis.SetupClient(&cfg.Redis)
	logZap := log.SetupLogger()
	log.Init(logZap)
	logger := log.GetLogger()

	userRepo := repositories.New(db, logger)
	userUsecase := usecases.New(userRepo, logger)
	middleware := middleware.Middleware{
		Repo: userRepo,
	}

	validator := validator.New()
	userHandler := handler.UserHandler{
		Log:       logger,
		Validator: validator,
		Usecase:   userUsecase,
	}

	serverHttp := http.SetupHttpEngine()

	r := router.Initialize(serverHttp, &userHandler, &middleware)

	return r

}
