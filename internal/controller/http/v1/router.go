package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/szymon676/jobguru/internal/usecase"
)

func SetupRoutes(j usecase.Job, u usecase.User) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	newJobRoutes(app, j)
	newUserRoutes(app, u, "1234")
	return app
}
