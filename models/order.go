package models

type Order struct {
	Id          uint      `json:"id"`
	FullName    string    `json:"fullname"`
	Address     string    `json:"address"`
	Note        string    `json:"note"`
	Phone       string    `json:"phone"`
	Status      uint      `json:"status"`
	TotalAmount uint      `json:"total_amount"`
	PaypalCode  string    `json:"paypal_code"`
	UserID      uint      `json:"user_id"`
	Products    []Product `gorm:"many2many:order_detail;" json:"products"`
}
