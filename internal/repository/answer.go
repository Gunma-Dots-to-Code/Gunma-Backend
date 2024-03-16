package repository

import (
	"gorm.io/gorm"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
)

type AnswerRepository interface {
	Create(answer *model.Answer) error
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{db}
}

type answerRepository struct {
	db *gorm.DB

}

func (ar *answerRepository) Create(answer *model.Answer) error {
	return ar.db.Create(answer).Error
}