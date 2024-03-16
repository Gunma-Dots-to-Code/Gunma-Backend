package repository

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(question *model.Question) error
	Get(id string) (*model.Question, error)
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db}
}

type questionRepository struct {
	db *gorm.DB
}

func (qr *questionRepository) Create(question *model.Question) error {
	return qr.db.Create(question).Error
}

func (qr *questionRepository) Get(id string) (*model.Question, error) {
	question := &model.Question{}
	if err := qr.db.First(question, id).Error; err != nil {
		return nil, err
	}
	return question, nil
}
