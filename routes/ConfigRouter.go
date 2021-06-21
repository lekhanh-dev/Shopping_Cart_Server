package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lekhanh-dev/Shopping_Cart_Server/controllers"
)

func ConfigAuthRouter(router *fiber.Router) {
	(*router).Get("/register", controllers.Register) //Liệt kê
}

// func ConfigUserRouter(router *fiber.Router) {
// 	(*router).Get("/", controllers.GetAllUser)
// 	(*router).Get("/:username", controllers.GetUserByUsername)
// 	(*router).Put("/:username", controllers.UpdateUserByUsername)
// 	(*router).Patch("/:username", controllers.UpdateInfoProductById)
// }

func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/", controllers.GetAllProduct)
	(*router).Get("/:id", controllers.GetProductById)
	(*router).Post("/", controllers.CreateProduct)
	(*router).Delete("/:id", controllers.DeleteProductById)
	(*router).Put("/", controllers.UpsertProduct)
	// (*router).Patch("/", controllers.UpdateProduct)
}

func ConfigCategoryRouter(router *fiber.Router) {
	(*router).Get("/", controllers.GetAllCategory)
	(*router).Get("/:id", controllers.GetCategoryById)
	(*router).Post("/", controllers.CreateCategory)
	(*router).Delete("/:id", controllers.DeleteCategoryById)
	(*router).Put("/", controllers.UpsertCategory)
	(*router).Patch("/", controllers.UpdateCategory)
}

func ConfigProductImageRouter(router *fiber.Router) {
	(*router).Get("/:id", controllers.GetImagesByProductId)
	(*router).Post("", controllers.CreateImagesForProduct)
	(*router).Delete(":id", controllers.DeleteImagesForProduct)
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/:id", controllers.GetReviewsByProductId)
	(*router).Post("", controllers.CreateReviewsForProduct)
	(*router).Delete(":id", controllers.DeleteReviewForProduct)
	(*router).Patch("", controllers.UpdateReview)
}
