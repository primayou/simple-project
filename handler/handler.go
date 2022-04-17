package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUsersHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func LoginHandler(c *fiber.Ctx) error {
	body := c.Body()
	fmt.Println(body)
	return c.SendString("Hello World!")
}

func ForgotPasswordHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func CreatePostHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func UpdatePostHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func DeletePostHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
