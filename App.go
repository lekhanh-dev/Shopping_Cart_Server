package main

import (
	"github.com/gofiber/fiber/v2"
	DB "github.com/lekhanh-dev/Shopping_Cart_Server/database"
	"github.com/lekhanh-dev/Shopping_Cart_Server/routes"
	"gorm.io/gorm"
)

type OK struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	DB.ConnectDB()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	authRouter := app.Group("/auth")
	routes.ConfigAuthRouter(&authRouter) //http://localhost:3000/api/book

	app.Listen(":3000")

}
