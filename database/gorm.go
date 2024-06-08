package database

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

func GetConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	connection.SetMaxIdleConns(15)
	connection.SetMaxOpenConns(50)
	connection.SetConnMaxLifetime(300 * time.Second)

	return db
}

// migrate create -ext sql -dir db/migrations create_users_table
// migrate -database "mysql://root:BA@rio2024@tcp(localhost:3306)/coptas" -path database/migrations up
