//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/controller"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/db"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/repository"
	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func Wire() *fiber.App {
	wire.Build(
		db.NewDB,
		router.NewRouter,
		controller.NewUserController,
		controller.NewCategoryController,
		controller.NewQuestionController,
		controller.NewAnswerController,
		repository.NewUserRepository,
		repository.NewCategoryRepository,
		repository.NewQuestionRepository,
		repository.NewAnswerRepository,
	)
	return &fiber.App{}
}
