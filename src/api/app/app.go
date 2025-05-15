package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/mikietechie/gopeople/api/handlers"
	"github.com/mikietechie/gopeople/api/middleware"
	"github.com/mikietechie/gopeople/config"
	"github.com/mikietechie/gopeople/docs"
)

// @title Go-people
// @version 1.0
// @description Go-people api

// @contact.name Mike Z
// @contact.email mzinyoni7@yandex.com

// @BasePath /api/v1
func New() *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		Prefork:       true,
		AppName:       "Go People",
	})
	log.Info(docs.SwaggerInfo)
	// middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New(logger.Config{
		Output: config.LOGGING_WRITER,
	}))
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/metrics", monitor.New())
	app.Use(middleware.Security)
	// routing
	app.Get("", handlers.Index)
	app.Route("/api/v1", func(router fiber.Router) {
		router.Get("/users", handlers.ReadUsers)
		router.Post("/users", handlers.CreateUser)
		router.Get("/users/:id", handlers.ReadUser)
		router.Put("/users/:id", handlers.UpdateUser)
		router.Patch("/users/:id", handlers.EditUser)
		router.Delete("/users/:id", handlers.DeleteUser)
	})
	// return app
	return app
}
