package models

type Category struct {
	Id           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	Image        string    `json:"image"`
	Description  string    `json:"description"`
	IsActive     bool      `json:"is_active"`
	Products     []Product `json:"products"`
}
