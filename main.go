package main

import (
	"errors"
	"fmt"
	"github.com/C241-PS120/bangkit-cloud-computing/database"
	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/C241-PS120/bangkit-cloud-computing/handler"
	"github.com/C241-PS120/bangkit-cloud-computing/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {

	var isProd bool
	if os.Getenv("ENVIRONMENT") == "production" {
		isProd = true
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Error(err)
			log.Fatal("Error loading .env file")
		}
	}

	db := database.GetConnection(isProd)

	articleRepository := repository.NewArticleRepository(db)
	articleHandler := handler.NewArticleHandler(articleRepository)

	app := fiber.New(
		fiber.Config{
			Prefork:      isProd,
			AppName:      "Coptas Article API",
			ErrorHandler: NewErrorHandler(),
		},
	)

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowOriginsFunc: func(origin string) bool {
			// Allow all origins from localhost or 127.0.0.1 with any port
			if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
				return true
			}
			return true
		},
	}))

	app.Get("/metrics", monitor.New())
	app.Use(healthcheck.New(
		healthcheck.Config{
			LivenessEndpoint:  "/live",
			ReadinessEndpoint: "/ready",
		}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	articles := v1.Group("/articles")
	articles.Get("/:id", articleHandler.GetArticleDetail)
	articles.Get("/", articleHandler.GetArticleList)
	articles.Post("/", articleHandler.CreateArticle)
	articles.Put("/:id", articleHandler.UpdateArticle)
	articles.Delete("/:id", articleHandler.DeleteArticle)

	label := v1.Group("/label")
	label.Get("/:label", articleHandler.GetArticleByLabel)

	app.Use(handler.NotFoundHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		return ctx.Status(code).JSON(dto.Envelope{
			Success: false,
			Error:   err.Error(),
		})
	}
}
