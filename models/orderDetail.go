package models

import (
	"time"
)

type OrderDetail struct {
	OrderID    uint      `gorm:"primaryKey" json:"order_id"`
	ProductID  uint      `gorm:"primaryKey" json:"product_id"`
	Quantity   uint      `json:"quantity"`
	Price      uint      `json:"price"`
	CreatedAt  time.Time `json:"create_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
