package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"github.com/pavel97go/gosuslugi/internal/apperr"
	"github.com/pavel97go/gosuslugi/internal/models"
	"github.com/pavel97go/gosuslugi/usecase"
)

type Handler struct{ uc usecase.ApplicationUsecase }

func New(uc usecase.ApplicationUsecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) Register(r fiber.Router) {
	v1 := r.Group("/v1")
	apps := v1.Group("/applications")

	apps.Post("/", h.createApplication)
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
		return toHTTPErr(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(createResp{ID: id})
}

func (h *Handler) getApplication(c *fiber.Ctx) error {
	id, ok := parseID(c)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	a, err := h.uc.GetByID(ctx, id)
	if err != nil {
		return toHTTPErr(c, err)
	}
	return c.JSON(a)
}

func (h *Handler) listApplications(c *fiber.Ctx) error {
	var f models.ApplicationFilter

	if v := c.Query("status"); v != "" {
		s := models.ApplicationStatus(v)
		f.Status = &s
	}
	if v := c.Query("document_type"); v != "" {
		t := models.ApplicationType(v)
		f.DocumentType = &t
	}
	if v := c.Query("limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			f.Limit = int32(n)
		}
	}
	if v := c.Query("offset"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			f.Offset = int32(n)
		}
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	items, err := h.uc.List(ctx, f)
	if err != nil {
		return toHTTPErr(c, err)
	}
	return c.JSON(items)
}

type updateReq struct {
	CitizenName  string                   `json:"citizen_name"`
	DocumentType models.ApplicationType   `json:"document_type"`
	Data         map[string]any           `json:"data"`
	Status       models.ApplicationStatus `json:"status"`
}

func (h *Handler) updateApplication(c *fiber.Ctx) error {
	id, ok := parseID(c)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req updateReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err := h.uc.Update(ctx, models.Application{
		ID:           id,
		CitizenName:  req.CitizenName,
		DocumentType: req.DocumentType,
		Data:         req.Data,
		Status:       req.Status,
	})
	if err != nil {
		return toHTTPErr(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) deleteApplication(c *fiber.Ctx) error {
	id, ok := parseID(c)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	if err := h.uc.Delete(ctx, id); err != nil {
		return toHTTPErr(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func toHTTPErr(c *fiber.Ctx, err error) error {
	log.Error().
		Err(err).
		Str("path", c.Path()).
		Str("method", c.Method()).
		Msg("request failed")

	switch err {
	case apperr.ErrValidation:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error"})
	case apperr.ErrNotFound:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	case apperr.ErrConflict:
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "conflict"})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
	}
}

func parseID(c *fiber.Ctx) (int64, bool) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	return id, err == nil && id > 0
}
