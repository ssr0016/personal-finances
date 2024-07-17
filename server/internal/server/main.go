package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/ssr0016/personal-finance/internal/config"
	"github.com/ssr0016/personal-finance/internal/controller"
	"github.com/ssr0016/personal-finance/internal/database"
	"github.com/ssr0016/personal-finance/internal/server/router"
	"github.com/ssr0016/personal-finance/internal/service"
)

type Server struct {
	app       *fiber.App
	port      string
	jwtSecret string
	db        *sqlx.DB
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: router.DefaultErrorHandler,
	})

	port := ":" + cfg.Port
	db := database.Connect(cfg.DatabaseUrl)

	return &Server{
		app:       app,
		port:      port,
		jwtSecret: cfg.JwtSecret,
		db:        db,
	}
}

func (s *Server) Start() error {
	us := service.NewUserService(s.db)

	uc := controller.NewAuthController(us, s.jwtSecret)

	s.SetupRoutes(uc)
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	s.db.Close()
	return s.app.Shutdown()
}
