package controller

import "github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/repository"

type QuestionController interface {
}

func NewQuestionController(r repository.QuestionRepository) QuestionController {
	return &questionController{r}
}

type questionController struct {
	questionRepository repository.QuestionRepository
}
