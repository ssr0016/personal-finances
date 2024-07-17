package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ssr0016/personal-finance/internal/model"
	"github.com/ssr0016/personal-finance/internal/server/router/response"
	"github.com/ssr0016/personal-finance/internal/service"
)

type TransactionController struct {
	s *service.TransactionService
}

func NewTransactionController(s *service.TransactionService) *TransactionController {
	return &TransactionController{
		s: s,
	}
}

func (c *TransactionController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	id := ctx.Params("id")
	transaction, err := c.s.GetById(id)

	if err != nil {
		return response.ErrorNotFound(err)
	}
	if transaction.UserId != userId {
		return response.ErrorUnauthorized(err, "unauthorized")
	}
	return response.Ok(ctx, fiber.Map{
		"transaction": transaction,
	})
}

func (c *TransactionController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	input := model.TransactionInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	transaction, err := c.s.Create(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Created(ctx, fiber.Map{
		"transaction": transaction,
	})
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	id := ctx.Params("id")
	input := model.TransactionInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	transaction, err := c.s.Update(id, userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"transaction": transaction,
	})
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	id := ctx.Params("id")
	transaction, err := c.s.Delete(id, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"transaction": transaction,
	})
}
