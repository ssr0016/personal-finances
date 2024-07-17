package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/ssr0016/personal-finance/internal/controller"
	"github.com/ssr0016/personal-finance/internal/server/router/middleware"
	"github.com/ssr0016/personal-finance/internal/server/router/response"
)

func healthCheck(db *sqlx.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var result int
		err := db.Get(&result, "SELECT 1")
		if err != nil {
			return errors.New("database unavailable")
		}
		return response.Ok(ctx, fiber.Map{
			"database": "available",
		})
	}
}

func (s *Server) SetupRoutes(uc *controller.AuthController) {
	api := s.app.Group("/api")
	api.Get("/", healthCheck(s.db))

	api.Post("/login", uc.Login)
	api.Post("/register", uc.Register)
	api.Get("/me", middleware.Authenticate(s.jwtSecret), uc.Me)
}
