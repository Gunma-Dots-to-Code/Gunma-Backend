package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	UserID        uint     `gorm:"not null"` //User削除時にAnswerも削除される
	QuestionID    uint     `gorm:"not null"`
	AnswerContent string   `gorm:"not null"`
	User          User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //User削除時にAnswerも削除される
	Question      Question `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
