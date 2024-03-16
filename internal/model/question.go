package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	UserID          uint     `gorm:"not null" json:"user_id"` //User削除時にQuestionも削除される
	CategoryID      uint     `gorm:"not null" json:"category_id"`
	QuestionTitle   string   `gorm:"not null" json:"question_title"`
	QuestionContent string   `gorm:"not null" json:"question_content"`
	User            User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"` //User削除時にQuestionも削除される
	Category        Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
}
