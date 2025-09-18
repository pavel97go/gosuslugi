package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pavel97go/gosuslugi/internal/apperr"
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

// GetByID
func (r *PostgresRepo) GetByID(ctx context.Context, id int64) (*models.Application, error) {
	const q = `
		SELECT id, citizen_name, document_type, data, status, created_at, updated_at
		FROM applications WHERE id = $1;
	`
	var a models.Application
	var dataRaw []byte
	if err := r.pool.QueryRow(ctx, q, id).Scan(
		&a.ID, &a.CitizenName, &a.DocumentType, &dataRaw, &a.Status, &a.CreatedAt, &a.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrNotFound
		}
		return nil, err
	}
	if len(dataRaw) > 0 {
		_ = json.Unmarshal(dataRaw, &a.Data)
	} else {
		a.Data = map[string]any{}
	}
	return &a, nil
}

// List
func (r *PostgresRepo) List(ctx context.Context, f models.ApplicationFilter) ([]models.Application, error) {
	var sb strings.Builder
	sb.WriteString(`
		SELECT id, citizen_name, document_type, data, status, created_at, updated_at
		FROM applications
	`)

	args := []any{}
	where := []string{}

	if f.Status != nil {
		where = append(where, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, *f.Status)
	}
	if f.DocumentType != nil {
		where = append(where, fmt.Sprintf("document_type = $%d", len(args)+1))
		args = append(args, *f.DocumentType)
	}
	if len(where) > 0 {
		sb.WriteString(" WHERE " + strings.Join(where, " AND "))
	}
	sb.WriteString(" ORDER BY created_at DESC ")

	limit := f.Limit
	if limit <= 0 || limit > 1000 {
		limit = 50
	}
	offset := f.Offset
	sb.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2))
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, sb.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Application
	for rows.Next() {
		var a models.Application
		var dataRaw []byte
		if err := rows.Scan(&a.ID, &a.CitizenName, &a.DocumentType, &dataRaw, &a.Status, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		if len(dataRaw) > 0 {
			_ = json.Unmarshal(dataRaw, &a.Data)
		}
		res = append(res, a)
	}
	return res, rows.Err()
}

// Update
func (r *PostgresRepo) Update(ctx context.Context, a *models.Application) error {
	const q = `
		UPDATE applications
		SET citizen_name = $1, document_type = $2, data = $3, status = $4
		WHERE id = $5;
	`
	payload, _ := json.Marshal(a.Data)
	ct, err := r.pool.Exec(ctx, q, a.CitizenName, a.DocumentType, payload, a.Status, a.ID)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return apperr.ErrNotFound
	}
	return nil
}

// UpdateStatus
func (r *PostgresRepo) UpdateStatus(ctx context.Context, id int64, newStatus models.ApplicationStatus) error {
	const q = `UPDATE applications SET status = $1 WHERE id = $2;`
	ct, err := r.pool.Exec(ctx, q, newStatus, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return apperr.ErrNotFound
	}
	return nil
}

// Delete
func (r *PostgresRepo) Delete(ctx context.Context, id int64) error {
	const q = `DELETE FROM applications WHERE id = $1;`
	ct, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return apperr.ErrNotFound
	}
	return nil
}
