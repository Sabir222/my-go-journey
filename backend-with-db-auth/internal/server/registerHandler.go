package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type userInfo struct {
	Full_name string `json:"full_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (s *FiberServer) RegisterHandler(c *fiber.Ctx) error {

	user := userInfo{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "something went wrong",
		})
	}

	fmt.Println(user.Email)

	return c.JSON(map[string]string{"message": "login"})
}
