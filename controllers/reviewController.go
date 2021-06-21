package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
	repo "github.com/lekhanh-dev/Shopping_Cart_Server/repository"
)

func GetReviewsByProductId(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	reviews := repo.GetRatingByProductId(uint(productId))
	return c.JSON(reviews)
}

func GetRatingByProductId(productId uint) float32 {
	reviews := repo.GetRatingByProductId(productId)
	var rating float32 = 0
	for _, review := range reviews {
		rating += float32(review.Point)
	}
	if rating == 0 {
		return 0
	} else {
		return rating / float32(len(reviews))
	}
}

func CreateReviewsForProduct(c *fiber.Ctx) error {
	var review models.Review
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	// check product exists
	product, isExists := repo.GetProductById(uint(review.ProductID))
	fmt.Println(product)
	if isExists {
		reviewId := repo.CreateReviewForProduct(&review)
		return c.SendString(fmt.Sprintf("New image is created successfully with id = %d", reviewId))
	} else {
		return c.Status(404).SendString("Product not exists")
	}
}

func DeleteReviewForProduct(c *fiber.Ctx) error {
	reviewId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	repo.DeleteReviewById(uint(reviewId))
	return c.SendString(fmt.Sprintf("Delete review with id = %d successfully", reviewId))
}

func UpdateReview(c *fiber.Ctx) error {
	var review models.Review
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.UpdateReview(&review)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Review with id = %d is successfully updated", review.Id))
}
