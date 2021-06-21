package models

type Cart struct {
	Id          int     `json:"id"`
	CategoryId  int     `json:"categoryId"`
	Image       string  `json:"image"`
	ProductName string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt  int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
