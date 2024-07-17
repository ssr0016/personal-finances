package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ssr0016/personal-finance/internal/model"
	"github.com/ssr0016/personal-finance/internal/server/router/response"
	"github.com/ssr0016/personal-finance/internal/service"
	"github.com/ssr0016/personal-finance/pkg/util"
)

type AuthController struct {
	s      *service.UserService
	secret string
}

func NewAuthController(s *service.UserService, secret string) *AuthController {
	return &AuthController{
		s:      s,
		secret: secret,
	}
}

func (c *AuthController) createToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.secret))
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}

	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	password, err := util.HashPassword(input.Password)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}
	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	token, err := c.createToken(user.Id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	return response.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}

	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "Login error")
	}
	if !util.CheckPassword(input.Password, user.Password) {
		return response.ErrorUnauthorized(err, "Login error")
	}

	token, err := c.createToken(user.Id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Login error")
	}

	return response.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)
	currentUser, err := c.s.GetByUsername(id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}

	token, err := c.createToken(currentUser.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}

	return response.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}
