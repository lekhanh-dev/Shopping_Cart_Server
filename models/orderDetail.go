package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	OrderID   uint           `gorm:"primaryKey" json:"order_id"`
	ProductID uint           `gorm:"primaryKey" json:"product_id"`
	Quantity  uint           `json:"quantity"`
	Price     uint           `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
