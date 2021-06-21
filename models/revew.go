package models

type Review struct {
	Id        uint    `json:"id"`
	Content   string  `json:"content"`
	Point     float32 `json:"point"`
	ProductID uint    `json:"product_id"`
	UserID    uint    `json:"user_id"`
}
