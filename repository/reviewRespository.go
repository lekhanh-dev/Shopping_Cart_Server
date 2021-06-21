package repository

import (
	"fmt"

	"github.com/lekhanh-dev/Shopping_Cart_Server/database"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
)

func GetRatingByProductId(productId uint) []models.Review {
	var reviews []models.Review
	database.DB.Where("product_id = (?)", productId).Find(&reviews)
	return reviews
}

func CreateReviewForProduct(review *models.Review) uint {
	database.DB.Create(&review)
	return review.Id
}

func DeleteReviewById(reviewId uint) {
	result := database.DB.Delete(&models.Review{}, reviewId)
	fmt.Println(result.Error)
}

func UpdateReview(review *models.Review) error {
	result := database.DB.Save(&review)
	return result.Error
}
