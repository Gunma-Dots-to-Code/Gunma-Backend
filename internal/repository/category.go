package repository

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) error
	List() ([]*model.Category, error)
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

func (cr *categoryRepository) List() ([]*model.Category, error) {
	var categories []*model.Category
	if err := cr.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
