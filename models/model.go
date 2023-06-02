package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primarykey"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Company struct {
	gorm.Model
	Name           string `json:"name"`
	IndustryType   string `json:"industry_type"`
	CompanyScaleId uint   `gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
}

type CompanyScale struct {
	gorm.Model
	Headquarter   string `json:"headquarter"`
	EmployeeTotal uint   `json:"employee_total"`
	Company       []*Company
}
