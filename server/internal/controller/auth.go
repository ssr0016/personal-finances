package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssr0016/personal-finance/internal/model"
	"github.com/ssr0016/personal-finance/internal/server/router"
	"github.com/ssr0016/personal-finance/internal/service"
	"github.com/ssr0016/personal-finance/pkg/util"
)

type AuthController struct {
	s *service.UserService
}

func NewAuthController(s *service.UserService) *AuthController {
	return &AuthController{
		s: s,
	}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}

	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	password, err := util.HashPassword(input.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "invalid password")
	}
	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "registration error")
	}
	return router.Created(ctx, fiber.Map{
		"user": user,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}

	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid username or password")
	}
	if !util.CheckPassword(input.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid username or password")
	}

	return router.Ok(ctx, fiber.Map{
		"user": user,
	})
}
