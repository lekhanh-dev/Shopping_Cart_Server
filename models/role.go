package models

type Role struct {
	Id       uint   `json:"id"`
	RoleName string `json:"role_name"`
	Users    []User `json:"users"`
}
