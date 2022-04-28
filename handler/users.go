package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetUser(c *fiber.Ctx) (err error) {
	ctx := c.Context()

	if err != nil {
		return err
	}

	users, _ := s.store.GetUsers(ctx)

	if err != nil {
		fmt.Println(err)
		return err
	}

	var response []*GetUserResponse

	for _, element := range users {
		response = append(response, &GetUserResponse{
			Email: element.Email,
		})
	}

	fmt.Println(response)

	return c.JSON(response)

}

type GetUserResponse struct {
	Email string `json:"email,omitempty"`
}
