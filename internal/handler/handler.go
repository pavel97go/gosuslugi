package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pavel97go/gosuslugi/internal/models"
	"github.com/pavel97go/gosuslugi/usecase"
)

type Handler struct{ uc usecase.ApplicationUsecase }

func New(uc usecase.ApplicationUsecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) Register(r fiber.Router) {
	v1 := r.Group("/v1")
	apps := v1.Group("/applications")

	apps.Post("/", h.createApplication)
	// Остальные пока заглушки
	apps.Get("/:id", h.getApplication)
	apps.Get("/", h.listApplications)
	apps.Put("/:id", h.updateApplication)
	apps.Delete("/:id", h.deleteApplication)
}

type createReq struct {
	CitizenName  string                 `json:"citizen_name"`
	DocumentType models.ApplicationType `json:"document_type"`
	Data         map[string]any         `json:"data"`
}

type createResp struct {
	ID int64 `json:"id"`
}

func (h *Handler) createApplication(c *fiber.Ctx) error {
	var req createReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}
	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	id, err := h.uc.Create(ctx, models.Application{
		CitizenName:  req.CitizenName,
		DocumentType: req.DocumentType,
		Data:         req.Data,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createResp{ID: id})
}

// Заглушки
func (h *Handler) getApplication(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNotImplemented) }
func (h *Handler) listApplications(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
func (h *Handler) updateApplication(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
func (h *Handler) deleteApplication(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
