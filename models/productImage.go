package models

import "time"

type ProductImage struct {
	Id          uint      `json:"id"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	ProductID   uint      `json:"product_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"create_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
