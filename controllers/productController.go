package controllers

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
	repo "github.com/lekhanh-dev/Shopping_Cart_Server/repository"
)

func GetAllProduct(c *fiber.Ctx) error {
	return c.JSON(repo.GetAllProduct())
}

func GetProductById(c *fiber.Ctx) error {
	idProduct, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product, isExists := repo.GetProductById(uint(idProduct))
	if isExists {
		productMap := structs.Map(product)
		productMap["rating"] = GetRatingByProductId(uint(idProduct))
		return c.JSON(productMap)
	} else {
		return c.Status(404).SendString("Product not exists")
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	// Check category exists
	categoryId := product.CategoryID
	category, isExists := repo.GetCategoryById(uint(categoryId))
	fmt.Println(category)
	if isExists {
		productId := repo.CreateProduct(&product)
		return c.SendString(fmt.Sprintf("New product is created successfully with id = %d", productId))
	} else {
		return c.Status(404).SendString("Category not exists")
	}
}

func DeleteProductById(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	repo.DeleteProductById(uint(productId))
	return c.SendString(fmt.Sprintf("Delete product with id = %d successfully", productId))
}

func UpsertProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	// Check category exists
	categoryId := product.CategoryID
	category, isExists := repo.GetCategoryById(uint(categoryId))
	fmt.Println(category)
	if isExists {
		productId := repo.UpsertProduct(&product)
		return c.SendString(fmt.Sprintf("Product with id = %d is successfully upserted", productId))
	} else {
		return c.Status(404).SendString("Category not exists")
	}

}

func UpdateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	err = repo.UpdateProduct(product)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", product.Id))
}
