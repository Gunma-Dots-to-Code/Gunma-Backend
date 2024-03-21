package controller

import (
	"strconv"

	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type QuestionController interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	ListByCategoryID(c *fiber.Ctx) error
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

	categoryID, err := strconv.Atoi(c.Params("category_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	question.CategoryID = uint(categoryID)
	if err := qc.questionRepository.Create(question); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(question)
}

func (qc *questionController) Get(c *fiber.Ctx) error {
	id := c.Params("question_id")

	//idはuint型だけどもstring型のままrepositoryにアクセスして問題ない
	question, err := qc.questionRepository.Get(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(question)
}

func (qc *questionController) ListByCategoryID(c *fiber.Ctx) error {
	// categoryに紐づくすべてのquestionを取得
	categoryID := c.Params("category_id")
	questions, err := qc.questionRepository.ListByCategoryID(categoryID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(questions)
}
