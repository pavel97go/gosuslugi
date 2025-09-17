package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pavel97go/gosuslugi/internal/models"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}

// Create: вставка заявки, возврат id
func (r *PostgresRepo) Create(ctx context.Context, a *models.Application) (int64, error) {
	const q = `
		INSERT INTO applications (citizen_name, document_type, data, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	row := r.pool.QueryRow(ctx, q, a.CitizenName, a.DocumentType, a.Data, a.Status)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// Заглушки (реализуем позже, чтобы компилировалось)
func (r *PostgresRepo) GetByID(ctx context.Context, id int64) (*models.Application, error) {
	return nil, errors.New("not implemented")
}
func (r *PostgresRepo) List(ctx context.Context, f models.ApplicationFilter) ([]models.Application, error) {
	return nil, errors.New("not implemented")
}
func (r *PostgresRepo) Update(ctx context.Context, a *models.Application) error {
	return errors.New("not implemented")
}
func (r *PostgresRepo) UpdateStatus(ctx context.Context, id int64, newStatus models.ApplicationStatus) error {
	return errors.New("not implemented")
}
func (r *PostgresRepo) Delete(ctx context.Context, id int64) error {
	return errors.New("not implemented")
}
