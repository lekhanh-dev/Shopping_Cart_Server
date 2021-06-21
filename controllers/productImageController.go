package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
	repo "github.com/lekhanh-dev/Shopping_Cart_Server/repository"
)

func GetImagesByProductId(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	productImages := repo.GetImagesByProductId(uint(productId))

	return c.JSON(productImages)
}

func CreateImagesForProduct(c *fiber.Ctx) error {
	var productImage models.ProductImage
	err := c.BodyParser(&productImage)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// check product exists
	product, isExists := repo.GetProductById(uint(productImage.ProductID))
	fmt.Println(product)
	if isExists {
		productImageId := repo.CreateImagesForProduct(&productImage)
		return c.SendString(fmt.Sprintf("New image is created successfully with id = %d", productImageId))
	} else {
		return c.Status(404).SendString("product not exists")
	}
}

func DeleteImagesForProduct(c *fiber.Ctx) error {
	productmageId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	repo.DeleteProductImageById(uint(productmageId))
	return c.SendString(fmt.Sprintf("Delete image with id = %d successfully", productmageId))
}
