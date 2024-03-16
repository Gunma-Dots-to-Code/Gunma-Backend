package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	UserID          uint     `gorm:"not null"` //User削除時にQuestionも削除される
	CategoryID      uint     `gorm:"not null"`
	QuestionTitle   string   `gorm:"not null"`
	QuestionContent string   `gorm:"not null"`
	User            User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //User削除時にQuestionも削除される
	Category        Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
