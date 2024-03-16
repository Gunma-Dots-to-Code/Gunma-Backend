package controller

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/model"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Create(c *fiber.Ctx) error
}

func NewCategoryController(r repository.CategoryRepository) CategoryController {
	return &categoryController{r}
}

type categoryController struct {
	categoryRepository repository.CategoryRepository
}

func (cc *categoryController) Create(c *fiber.Ctx) error {
	category := &model.Category{}
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := cc.categoryRepository.Create(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}