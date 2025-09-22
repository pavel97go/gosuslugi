package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pavel97go/gosuslugi/internal/handler"
	"github.com/pavel97go/gosuslugi/internal/repository"
	"github.com/pavel97go/gosuslugi/usecase"
)

func NewRouter(pool *pgxpool.Pool) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health/live", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/health/ready", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "up",
			"version": "v1.0.1",
		})
	})

	repo := repository.NewPostgresRepo(pool)
	uc := usecase.New(repo)
	h := handler.New(uc)
	h.Register(app)

	return app
}
