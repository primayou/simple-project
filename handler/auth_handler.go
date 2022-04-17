package handler

import (
	"encoding/json"
	"fmt"
	"simple-projects/repository"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) RegistrationHandler(c *fiber.Ctx) (err error) {

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

	err = server.store.Registration(c.Context(), arg)

	if err != nil {
		fmt.Print(err)
	}

	return c.JSON(incomingRequest)
}

func (server *Server) LoginHandler(c *fiber.Ctx) (err error) {
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
