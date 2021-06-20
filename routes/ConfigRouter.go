package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/lekhanh-dev/Shopping_Cart_Server/controllers"
)

func ConfigAuthRouter(router *fiber.Router) {
	(*router).Get("/register", controller.Register) //Liệt kê

}
