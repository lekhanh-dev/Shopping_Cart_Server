package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/lekhanh-dev/Shopping_Cart_Server/database"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
)

func GetAllProduct() []models.Product {
	var products []models.Product
	database.DB.Find(&products)
	return products
}

func GetProductById(idProduct uint) (models.Product, bool) {
	var product models.Product
	database.DB.Where("id = (?)", idProduct).Find(&product)
	if product.Id == 0 {
		return product, false
	} else {
		return product, true
	}
}

func CreateProduct(product *models.Product) uint {
	database.DB.Create(product)
	return product.Id
}

func DeleteProductById(productId uint) {
	database.DB.Where("product_id = (?)", productId).Delete(models.ProductImage{})
	database.DB.Where("product_id = (?)", productId).Delete(models.Review{})
	result := database.DB.Delete(&models.Product{}, productId)
	fmt.Println(result.Error)
}

func UpsertProduct(product *models.Product) uint {
	productById, isExists := GetProductById(product.Id)
	if isExists {
		productById.ProductName = product.ProductName
		productById.CategoryID = product.CategoryID
		productById.Image = product.Image
		productById.NumberAvailable = product.NumberAvailable
		productById.Price = product.Price
		productById.IsActive = product.IsActive
		productById.SalePercent = product.SalePercent
		productById.ModifiedAt = time.Now()
		database.DB.Save(&productById)
	} else {
		product.Id = 0
		CreateProduct(product)
	}
	return product.Id
}

func UpdateProduct(product *models.Product) error {
	productById, isExists := GetProductById(product.Id)
	if isExists {
		productById = *product
		database.DB.Save(&productById)
		return nil
	} else {
		return errors.New("Product not found")
	}
}
