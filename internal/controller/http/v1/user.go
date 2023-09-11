package v1

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/jobguru/internal/entity"
	"github.com/szymon676/jobguru/internal/usecase"
)

type UserRoutes struct {
	UseCase usecase.User
	Secret  []byte
}

func newUserRoutes(router *fiber.App, useCase usecase.User, secretKey string) {
	userRoutes := UserRoutes{
		UseCase: useCase,
		Secret:  []byte(secretKey),
	}

	authGroup := router.Group("/auth")
	authGroup.Post("/register", userRoutes.RegisterUser)
	authGroup.Post("/login", userRoutes.LoginUser)
	authGroup.Post("/logout", userRoutes.LogoutUser)
}

func (ur *UserRoutes) generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(ur.Secret)
}

func (ur *UserRoutes) RegisterUser(c *fiber.Ctx) error {
	var req entity.RegisterUser
	c.BodyParser(&req)

	if err := ur.UseCase.CreateUser(req); err != nil {
		return err
	}

	return c.JSON("User registration done successfully")
}

func (ur *UserRoutes) LoginUser(c *fiber.Ctx) error {
	var req entity.LoginUser
	c.BodyParser(&req)

	userID, err := ur.UseCase.LoginUser(req)
	if err != nil {
		return err
	}

	token, err := ur.generateToken(userID)
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:    "jwt_token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "user logged in successfully"})
}

func (ur *UserRoutes) LogoutUser(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:    "jwt_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "user logged out successfully"})
}
