package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"not null" json:"category_name"`
	// ParentCategoryID uint
}
