package repository

import (
	"errors"
	"fmt"

	"github.com/lekhanh-dev/Shopping_Cart_Server/database"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
)

func GetAllCategory() []models.Category {
	var categories []models.Category
	database.DB.Find(&categories)
	return categories
}

func GetCategoryById(categoryId uint) (models.Category, bool) {
	var category models.Category
	database.DB.Where("id = (?)", categoryId).Find(&category)
	if category.Id == 0 {
		return category, false
	} else {
		return category, true
	}
}

func CreateCategory(category *models.Category) uint {
	database.DB.Create(category)
	return category.Id
}

func DeleteCategoryById(categoryId uint) {
	result := database.DB.Delete(&models.Category{}, categoryId)
	fmt.Println(result.Error)
}

func UpdateCategory(category *models.Category) error {
	categoryById, isExists := GetCategoryById(category.Id)
	if isExists {
		categoryById = *category
		database.DB.Save(&categoryById)
		return nil
	} else {
		return errors.New("Category not found")
	}
}

func UpsertCategory(category *models.Category) uint {
	categoryById, isExists := GetCategoryById(category.Id)
	if isExists {
		categoryById.CategoryName = category.CategoryName
		database.DB.Save(&categoryById)
	} else {
		category.Id = 0
		CreateCategory(category)
	}
	return category.Id
}
