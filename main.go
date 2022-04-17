package main

import (
	"log"
	"simple-projects/handler"
	"simple-projects/repository"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"

	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := getDBSourceURL()

	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dbURL)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := repository.NewStore(conn)

	application := handler.NewServer(app, store)

	// Get Config
	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Open Port at " + config.AppConfig.AppPort)
	application.Listen(":" + config.AppConfig.AppPort)
}

func getDBSourceURL() string {

	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
}
