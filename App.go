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
	routes.ConfigAuthRouter(&authRouter)

	productRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&productRouter) //http://localhost:3000/api/product

	productImageRouter := app.Group("/api/product/image")
	routes.ConfigProductImageRouter(&productImageRouter)

	reviewRouter := app.Group("/api/review")
	routes.ConfigReviewRouter(&reviewRouter)

	categoryRouter := app.Group("/api/category")
	routes.ConfigCategoryRouter(&categoryRouter)

	app.Listen(":3000")
}
