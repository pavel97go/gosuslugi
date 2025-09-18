package usecase

import (
	"context"

	"github.com/pavel97go/gosuslugi/internal/apperr"
	"github.com/pavel97go/gosuslugi/internal/models"
	"github.com/pavel97go/gosuslugi/internal/repository"
)

type appUsecase struct {
	repo repository.ApplicationRepository
}

func New(repo repository.ApplicationRepository) ApplicationUsecase { return &appUsecase{repo: repo} }

// Create уже был, оставим:
func (u *appUsecase) Create(ctx context.Context, a models.Application) (int64, error) {
	if a.CitizenName == "" {
		return 0, apperr.ErrValidation
	}
	switch a.DocumentType {
	case models.TypePassport, models.TypeCertificate:
	default:
		return 0, apperr.ErrValidation
	}
	if a.Status == "" {
		a.Status = models.StatusDraft
	}
	return u.repo.Create(ctx, &a)
}

func (u *appUsecase) GetByID(ctx context.Context, id int64) (*models.Application, error) {
	if id <= 0 {
		return nil, apperr.ErrValidation
	}
	return u.repo.GetByID(ctx, id)
}

func (u *appUsecase) List(ctx context.Context, f models.ApplicationFilter) ([]models.Application, error) {
	if f.Limit <= 0 || f.Limit > 1000 {
		f.Limit = 50
	}
	if f.Offset < 0 {
		f.Offset = 0
	}
	return u.repo.List(ctx, f)
}

func (u *appUsecase) Update(ctx context.Context, a models.Application) error {
	if a.ID <= 0 || a.CitizenName == "" {
		return apperr.ErrValidation
	}
	switch a.DocumentType {
	case models.TypePassport, models.TypeCertificate:
	default:
		return apperr.ErrValidation
	}
	switch a.Status {
	case models.StatusDraft, models.StatusSubmitted, models.StatusApproved, models.StatusRejected:
	default:
		return apperr.ErrValidation
	}
	return u.repo.Update(ctx, &a)
}

func (u *appUsecase) UpdateStatus(ctx context.Context, id int64, newStatus models.ApplicationStatus) error {
	if id <= 0 {
		return apperr.ErrValidation
	}
	switch newStatus {
	case models.StatusDraft, models.StatusSubmitted, models.StatusApproved, models.StatusRejected:
	default:
		return apperr.ErrValidation
	}
	return u.repo.UpdateStatus(ctx, id, newStatus)
}

func (u *appUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return apperr.ErrValidation
	}
	return u.repo.Delete(ctx, id)
}
