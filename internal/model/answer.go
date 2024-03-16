package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	UserID        uint     `gorm:"not null" json:"user_id"` //User削除時にAnswerも削除される
	QuestionID    uint     `gorm:"not null" json:"question_id"`
	AnswerContent string   `gorm:"not null" json:"answer_content"`
	User          User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"` //User削除時にAnswerも削除される
	Question      Question `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"question"`
}
