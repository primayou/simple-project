package handler

import (
	"encoding/json"
	"fmt"
	"simple-projects/models"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) RegistrationHandler(c *fiber.Ctx) (err error) {
	ctx := c.Context()
	body := c.Body()

	incomingRequest := &RegistrationRequest{}
	err = json.Unmarshal(body, incomingRequest)

	if err != nil {
		return err
	}

	isUserExist, _ := s.store.GetUser(ctx, incomingRequest.Email)

	if isUserExist != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "something wrong")
	}

	mUser := &models.User{}
	mUser.Email = incomingRequest.Email
	mUser.PasswordHash = encrypt(incomingRequest.Password)

	err = s.store.CreateUser(ctx, mUser)

	if err != nil {
		return err
	}

	return c.JSON(&RegistrationResponse{
		Message: "success",
	})
}

func (s *Server) LoginHandler(c *fiber.Ctx) (err error) {
	ctx := c.Context()
	body := c.Body()

	incomingRequest := &LoginRequest{}
	err = json.Unmarshal(body, incomingRequest)

	if err != nil {
		fmt.Println(err)
	}

	user, err := s.store.GetUser(ctx, incomingRequest.Email)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Internal Server Error")
	}

	if user == nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "user not found")
	}

	if user.PasswordHash != encrypt(incomingRequest.Password) {
		return fiber.NewError(fiber.ErrBadRequest.Code, "credential is not valid")
	}

	token, err := CreateAccessToken(user)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Internal Server Error")
	}

	return c.JSON(&LoginResponse{
		Token: token,
	})
}

type RegistrationRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
type RegistrationResponse struct {
	Message string `json:"message,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
