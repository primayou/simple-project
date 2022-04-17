package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	application := InitRouter(app)

	config, err := getConfig()

	if err != nil {
		panic(err)
	}

	fmt.Println("Open Port at " + config.AppConfig.AppPort)
	application.Listen(":" + config.AppConfig.AppPort)
}
