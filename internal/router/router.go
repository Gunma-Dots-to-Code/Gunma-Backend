package router

import (
	"os"

	"github.com/Gunma-Dots-to-Code/Gunma-Backend/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewRouter(
	userController controller.UserController,
	categoryController controller.CategoryController,
	questionController controller.QuestionController,
	answerController controller.AnswerController,
) *fiber.App {
	app := setupRouter()

	//health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRouter := app.Group("/users")
	{
		userRouter.Post("/", userController.Create)
	}

	categoryRouter := app.Group("/categories")
	{
		categoryRouter.Post("/", categoryController.Create)
		categoryRouter.Get("/", categoryController.List)
	}

	questionRouter := app.Group("/questions")
	{
		questionRouter.Post("/", questionController.Create)
		questionRouter.Get("/", questionController.List)
		questionRouter.Get("/:id", questionController.Get)
	}

	answerRouter := app.Group("/answers")
	{
		answerRouter.Post("/", answerController.Create)
	}

	return app
}

func setupRouter() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOW_ORIGIN"),
		AllowMethods: "GET, POST, PATCH, DELETE, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Authorization",
	}))

	return app
}
