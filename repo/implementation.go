package repo

import "gorm.io/gorm"

type Implementation struct {
	db *gorm.DB
}

func NewImplementation(db *gorm.DB) *Implementation {
	return &Implementation{
		db: db,
	}
}
