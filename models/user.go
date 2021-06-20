package models

type User struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `gorm:"unique" json:"email"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Avatar   string  `json:"avatar"`
	IsActive bool    `json:"is_active"`
	RoleID   uint    `json:"role_id"`
	Orders   []Order `json:"orders"`
}
