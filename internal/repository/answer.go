package repository

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(answer *model.Answer) error
	GetByQuestionID(id string) ([]*model.Answer, error)
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

func (ar *answerRepository) GetByQuestionID(id string) ([]*model.Answer, error) {
	var answers []*model.Answer
	if err := ar.db.Preload("User").Preload("Question").Where("question_id = ?", id).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}