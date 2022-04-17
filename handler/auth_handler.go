package handler

import (
	"encoding/json"
	"fmt"
	"simple-projects/repository"

	"github.com/gofiber/fiber/v2"
)

func RegistrationHandler(c *fiber.Ctx) (err error) {
	body := c.Body()

	incomingRequest := &RegistrationRequest{}
	err = json.Unmarshal(body, incomingRequest)

	if err != nil {
		fmt.Println(err)
	}

	arg := repository.RegistrationParams{
		Email:        incomingRequest.Email,
		PasswordHash: incomingRequest.Password,
	}

	fmt.Println(arg)

	return c.JSON(incomingRequest)
}

func LoginHandler(c *fiber.Ctx) (err error) {
	body := c.Body()

	incomingRequest := &LoginRequest{}
	err = json.Unmarshal(body, incomingRequest)

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(incomingRequest)
}

type RegistrationRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
