package repository

import (
	"fmt"

	"github.com/lekhanh-dev/Shopping_Cart_Server/database"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
)

func GetImagesByProductId(productId uint) []models.ProductImage {
	var productImages []models.ProductImage
	database.DB.Where("product_id = (?)", productId).Find(&productImages)
	return productImages
}

func CreateImagesForProduct(productImage *models.ProductImage) uint {
	database.DB.Create(productImage)
	return productImage.Id
}

func DeleteProductImageById(productmageId uint) {
	result := database.DB.Delete(&models.ProductImage{}, productmageId)
	fmt.Println(result.Error)
}
