package models

type User struct {
	ID       uint   `gorm:"primarykey"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
