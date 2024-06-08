package main

import (
	"errors"
	"fmt"
	"github.com/C241-PS120/bangkit-cloud-computing/database"
	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/C241-PS120/bangkit-cloud-computing/handler"
	"github.com/C241-PS120/bangkit-cloud-computing/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	log.Println(os.Getenv("ENVIRONMENT"))
	log.Println(os.Getenv("PORT"))
	log.Println(os.Getenv("DB_NAME"))
	log.Println(os.Getenv("ENVIRONMENT") == "production")
	var isProd bool
	if os.Getenv("ENVIRONMENT") == "production" {
		isProd = true
	} else {
		err := godotenv.Load()
		if err != nil {
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

	// under development
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins:     "*",
	//	AllowCredentials: true,
	//}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	article := v1.Group("/articles")
	article.Get("/:id", articleHandler.GetArticleDetail)
	article.Get("/", articleHandler.GetArticleList)

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
