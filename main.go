package main

import (
	"flag"
	"github.com/C241-PS120/bangkit-cloud-computing/database"
	"github.com/C241-PS120/bangkit-cloud-computing/handler"
	"github.com/C241-PS120/bangkit-cloud-computing/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"

	"github.com/joho/godotenv"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	db := database.GetConnection()
	//err = db.AutoMigrate(&model.Article{}, &model.Category{}, &model.Symptom{}, &model.Prevention{}, &model.Treatment{})
	//if err != nil {
	//	log.Fatal(err)
	//}

	articleRepository := repository.NewArticleRepository(db)
	articleHandler := handler.NewArticleHandler(articleRepository)

	flag.Parse()

	app := fiber.New(
		fiber.Config{
			Prefork: *prod,
		},
	)
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/:id", articleHandler.GetArticleDetail)
	app.Get("/", articleHandler.GetArticleList)

	// Listen to the port
	log.Fatal(app.Listen(*port))
}
