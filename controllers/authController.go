package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		Name:     "Hello world",
		Email:    "hello@gmail.com",
		Password: "XXXXXXXXXXX",
		Phone:    "0331234569",
		IsActive: true,
		RoleID:   1,
	}

	// var user map[string]int
	// user["name"] = 12
	// user["age"] = 23

	// err := c.BodyParser(&user)
	// if err != nil {
	// 	return c.JSON(err)
	// }

	return c.JSON(user)
}
