package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
	repo "github.com/lekhanh-dev/Shopping_Cart_Server/repository"
)

func GetAllCategory(c *fiber.Ctx) error {
	return c.JSON(repo.GetAllCategory())
}

func GetCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	category, isExists := repo.GetCategoryById(uint(categoryId))
	if isExists {
		return c.JSON(category)
	} else {
		return c.Status(404).SendString("Category not exists")
	}
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	categoryId := repo.CreateCategory(&category)

	return c.SendString(fmt.Sprintf("New category is created successfully with id = %d", categoryId))
}

func DeleteCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	repo.DeleteCategoryById(uint(categoryId))
	return c.SendString(fmt.Sprintf("Delete category with id = %d successfully", categoryId))
}

func UpdateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	err = repo.UpdateCategory(category)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Category with id = %d is successfully updated", category.Id))
}

func UpsertCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	categoryId := repo.UpsertCategory(category)
	return c.SendString(fmt.Sprintf("Category with id = %d is successfully upserted", categoryId))
}
