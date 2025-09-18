package app

import (
	"context"
	"time"

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

	app.Get("/health/live", func(c *fiber.Ctx) error { return c.SendString("ok") })
	app.Get("/health/ready", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
		defer cancel()
		if err := pool.Ping(ctx); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "down", "error": err.Error()})
		}
		return c.JSON(fiber.Map{"status": "up"})
	})

	repo := repository.NewPostgresRepo(pool)
	uc := usecase.New(repo)
	h := handler.New(uc)
	h.Register(app)

	return app
}
