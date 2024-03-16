package repository

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

type categoryRepository struct {
	db *gorm.DB
}

func (cr *categoryRepository) Create(category *model.Category) error {
	return cr.db.Create(category).Error
}