package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/ssr0016/personal-finance/internal/config"
	"github.com/ssr0016/personal-finance/internal/database"
	"github.com/ssr0016/personal-finance/internal/server/router"
)

type Server struct {
	app  *fiber.App
	port string
	db   *sqlx.DB
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: router.DefaultErrorHandler,
	})
	port := ":" + cfg.Port
	db := database.Connect(cfg.DatabaseUrl)

	return &Server{
		app:  app,
		port: port,
		db:   db,
	}
}

func (s *Server) Start() error {
	s.SetupRoutes()
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	s.db.Close()
	return s.app.Shutdown()
}
