package repository

import (
	"context"

	"github.com/pavel97go/gosuslugi/internal/models"
)

type ApplicationRepository interface {
	// Создать заявку, вернуть ID
	Create(ctx context.Context, a *models.Application) (int64, error)

	// Получить по ID
	GetByID(ctx context.Context, id int64) (*models.Application, error)

	// Список с фильтрами и пагинацией
	List(ctx context.Context, f models.ApplicationFilter) ([]models.Application, error)

	// Полное обновление (id в a.ID обязателен)
	Update(ctx context.Context, a *models.Application) error

	// Изменить статус (с проверками на уровне usecase)
	UpdateStatus(ctx context.Context, id int64, newStatus models.ApplicationStatus) error

	// Удалить по ID
	Delete(ctx context.Context, id int64) error
}
