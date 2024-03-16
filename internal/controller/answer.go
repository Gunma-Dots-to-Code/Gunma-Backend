package controller

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type AnswerController interface {
	Create(c *fiber.Ctx) error
	GetByQuestionID(c *fiber.Ctx) error
}

func NewAnswerController(r repository.AnswerRepository) AnswerController {
	return &answerController{r}
}

type answerController struct {
	answerRepository repository.AnswerRepository
}

func (ac *answerController) Create(c *fiber.Ctx) error {
	answer := &model.Answer{}
	if err := c.BodyParser(answer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ac.answerRepository.Create(answer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(answer)
}


func (ac *answerController) GetByQuestionID(c *fiber.Ctx) error {
	question_id := c.Params("question_id")

	answer, err := ac.answerRepository.GetByQuestionID(question_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(answer)
}
