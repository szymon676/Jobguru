package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/jobguru/internal/usecase"
)

func SetupRoutes(j usecase.Job, u usecase.User) *fiber.App {
	app := fiber.New()
	newJobRoutes(app, j)
	newUserRoutes(app, u)
	return app
}
