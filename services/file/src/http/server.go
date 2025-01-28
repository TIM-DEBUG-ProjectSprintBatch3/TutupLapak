package httpServer

import (
	"fmt"
	"time"

	"github.com/TimDebug/TutupLapak/File/src/config"
	"github.com/TimDebug/TutupLapak/File/src/database/postgres"
	"github.com/TimDebug/TutupLapak/File/src/http/middleware/errorHandler"
	"github.com/TimDebug/TutupLapak/File/src/http/middleware/identifier"
	localLog "github.com/TimDebug/TutupLapak/File/src/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HttpServer struct{}

func (s *HttpServer) Listen() {
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
		ErrorHandler: errorHandler.ErrorHandler,
	})
	app.Use(identifier.RequestID)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST",
	}))
	app.Use(logger.New(logger.Config{
		Done:          localLog.ZerologWriter,
		TimeFormat:    time.RFC3339Nano,
		DisableColors: true,
	}))

	db, err := postgres.NewPgxConnect()
	if err != nil {
		localLog.Logger.Fatal().Err(err).Msg("unable to establish database connection")
	}
	appConfig := config.Config
	var storageClient StorageClient
	if appConfig.IsProduction {
		storageClient = NewS3StorageClient()
	} else {
		storageClient = NewMockS3StorageClient()
	}
	repo := NewFileRepository(db)
	service := NewFileService(repo, storageClient)
	defer service.Shutdown()
	controller := NewFileController(service)

	routes := app.Group("/v1")
	routes.Post("/file", controller.Upload)
	routes.Get("/file/:fileId", controller.CheckExist)

	app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort()))
}
