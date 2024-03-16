package repository

import "gorm.io/gorm"

type QuestionRepository interface {
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db}
}

type questionRepository struct {
	db *gorm.DB
}
