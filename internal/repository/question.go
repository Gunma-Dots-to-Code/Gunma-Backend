package repository

import(
	"gorm.io/gorm"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
)

type QuestionRepository interface {
	Create(question *model.Question) error
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