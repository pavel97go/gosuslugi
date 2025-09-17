package usecase

import (
	"context"
	"errors"

	"github.com/pavel97go/gosuslugi/internal/models"
	"github.com/pavel97go/gosuslugi/internal/repository"
)

type appUsecase struct {
	repo repository.ApplicationRepository
}

func New(repo repository.ApplicationRepository) ApplicationUsecase {
	return &appUsecase{repo: repo}
}

func (u *appUsecase) Create(ctx context.Context, a models.Application) (int64, error) {
	// простая валидация
	if a.CitizenName == "" {
		return 0, errors.New("citizen_name is required")
	}
	switch a.DocumentType {
	case models.TypePassport, models.TypeCertificate:
	default:
		return 0, errors.New("document_type must be 'passport' or 'certificate'")
	}
	// статус по умолчанию
	if a.Status == "" {
		a.Status = models.StatusDraft
	}
	return u.repo.Create(ctx, &a)
}

// Заглушки для остальных методов
func (u *appUsecase) GetByID(ctx context.Context, id int64) (*models.Application, error) {
	return nil, errors.New("not implemented")
}
func (u *appUsecase) List(ctx context.Context, f models.ApplicationFilter) ([]models.Application, error) {
	return nil, errors.New("not implemented")
}
func (u *appUsecase) Update(ctx context.Context, a models.Application) error {
	return errors.New("not implemented")
}
func (u *appUsecase) UpdateStatus(ctx context.Context, id int64, newStatus models.ApplicationStatus) error {
	return errors.New("not implemented")
}
func (u *appUsecase) Delete(ctx context.Context, id int64) error {
	return errors.New("not implemented")
}
