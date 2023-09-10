package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/jobguru/internal/entity"
	"github.com/szymon676/jobguru/internal/usecase"
)

type userRoutes struct {
	usecase usecase.User
}

func newUserRoutes(router *fiber.App, usecase usecase.User) {
	r := userRoutes{usecase: usecase}

	ur := router.Group("/auth")
	ur.Post("/register", (r.registerUser))
	ur.Post("/login", (r.loginUser))
	ur.Get("/users/{id}", (r.getUserByID))
}

func (uh *userRoutes) registerUser(c *fiber.Ctx) error {
	var req entity.RegisterUser

	c.BodyParser(&req)

	if err := uh.usecase.CreateUser(req); err != nil {
		return err
	}

	return c.JSON("User registration done successfully")
}

func (uh *userRoutes) loginUser(c *fiber.Ctx) error {
	var req entity.LoginUser

	c.BodyParser(&req)

	token, err := uh.usecase.LoginUser(req)
	if err != nil {
		return err
	}

	_ = token

	return c.JSON("user logged in successfully")
}

func (uh *userRoutes) getUserByID(c *fiber.Ctx) error {
	userid := c.Params("id")
	id, _ := strconv.Atoi(userid)

	user, err := uh.usecase.GetUserByID(id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
