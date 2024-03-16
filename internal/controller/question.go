package controller

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type QuestionController interface {
	Create(c *fiber.Ctx) error
}

func NewQuestionController(r repository.QuestionRepository) QuestionController {
	return &questionController{r}
}

type questionController struct {
	questionRepository repository.QuestionRepository
}

func (qc *questionController) Create(c *fiber.Ctx) error {
	question := &model.Question{}
	if err := c.BodyParser(question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := qc.questionRepository.Create(question); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(question)
}
