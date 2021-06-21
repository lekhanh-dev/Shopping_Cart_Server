package models

import "time"

type User struct {
	Id         uint       `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Username   string     `gorm:"unique" json:"username"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Gender     string     `json:"gender"`
	Phone      string     `json:"phone"`
	Birthday   *time.Time `json:"birthday"`
	Avatar     string     `json:"avatar"`
	IsActive   bool       `json:"status"`
	CreatedAt  time.Time  `json:"create_at"`
	ModifiedAt time.Time  `json:"modified_at"`
	RoleID     uint       `json:"role_id"`
	Orders     []Order    `json:"orders"`
	Reviews    []Review   `json:"reviews"`
}
