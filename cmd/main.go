package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"user-service/config"
	"user-service/internal/module/user/handler"
	"user-service/internal/module/user/repositories"
	"user-service/internal/module/user/usecases"
	"user-service/internal/pkg/database"
	"user-service/internal/pkg/helpers/middleware"
	"user-service/internal/pkg/http"
	log_internal "user-service/internal/pkg/log"
	router "user-service/internal/route"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
)

func main() {
	cfg := config.InitConfig()

	app := initService(cfg)

	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		log.Print("Starting host instrumentation:")
		err = host.Start(host.WithMeterProvider(provider))
		if err != nil {
			log.Fatal(err)
		}

		<-ctx.Done()
	}()

	// start http server
	http.StartHttpServer(app, cfg.HttpServer.Port)
}

func initService(cfg *config.Config) *fiber.App {
	db := database.GetConnection(&cfg.Database)
	// redis := redis.SetupClient(&cfg.Redis)
	logger := log_internal.Setup()

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
	ctx := context.Background()
	conn, serviceName, err := http.InitConn(cfg)
	if err != nil {
		logger.Ctx(ctx).Fatal(fmt.Sprintf("Failed to create gRPC connection to collector: %v", err))
	}

	// setup tracer
	http.InitTracer(conn, serviceName)

	// setup metric
	_, err = http.InitMeterProvider(conn, serviceName)
	if err != nil {
		logger.Ctx(ctx).Fatal(fmt.Sprintf("Failed to create meter provider: %v", err))
	}

	r := router.Initialize(serverHttp, &userHandler, &middleware)

	return r

}
