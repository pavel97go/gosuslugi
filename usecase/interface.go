package usecase

import (
	"context"

	"github.com/pavel97go/gosuslugi/internal/models"
)

type ApplicationUsecase interface {
	Create(ctx context.Context, a models.Application) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Application, error)
	List(ctx context.Context, f models.ApplicationFilter) ([]models.Application, error)
	Update(ctx context.Context, a models.Application) error
	UpdateStatus(ctx context.Context, id int64, newStatus models.ApplicationStatus) error
	Delete(ctx context.Context, id int64) error
}
