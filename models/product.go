package models

import "time"

type Product struct {
	Id              uint   `json:"id"`
	ProductName     string `json:"product_name"`
	Image           string `json:"image"`
	NumberAvailable uint   `json:"number_available"`
	Price           uint   `json:"price"`
	Description     string `json:"description"`
	Properties      string `json:"properties"`
	IsActive        bool   `json:"is_active"`
	SalePercent     uint   `json:"sale_percent"`
	CategoryID      uint   `json:"category_id"`
	ProductImages   []ProductImage
	Reviews         []Review  `json:"reviews"`
	CreatedAt       time.Time `json:"create_at"`
	ModifiedAt      time.Time `json:"modified_at"`
}
