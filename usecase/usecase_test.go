package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/pavel97go/gosuslugi/internal/apperr"
	"github.com/pavel97go/gosuslugi/internal/models"
	"github.com/pavel97go/gosuslugi/internal/repository"
)

type fakeRepo struct{ items []models.Application }

func (f *fakeRepo) Create(ctx context.Context, a *models.Application) (int64, error) {
	a.ID = int64(len(f.items) + 1)
	now := time.Now()
	a.CreatedAt, a.UpdatedAt = now, now
	f.items = append(f.items, *a)
	return a.ID, nil
}
func (f *fakeRepo) GetByID(ctx context.Context, id int64) (*models.Application, error) {
	for i := range f.items {
		if f.items[i].ID == id {
			cp := f.items[i]
			return &cp, nil
		}
	}
	return nil, apperr.ErrNotFound
}
func (f *fakeRepo) List(ctx context.Context, _ models.ApplicationFilter) ([]models.Application, error) {
	out := make([]models.Application, len(f.items))
	copy(out, f.items)
	return out, nil
}
func (f *fakeRepo) Update(ctx context.Context, a *models.Application) error {
	for i := range f.items {
		if f.items[i].ID == a.ID {
			a.CreatedAt = f.items[i].CreatedAt
			a.UpdatedAt = time.Now()
			f.items[i] = *a
			return nil
		}
	}
	return apperr.ErrNotFound
}
func (f *fakeRepo) UpdateStatus(ctx context.Context, id int64, s models.ApplicationStatus) error {
	for i := range f.items {
		if f.items[i].ID == id {
			f.items[i].Status = s
			f.items[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return apperr.ErrNotFound
}
func (f *fakeRepo) Delete(ctx context.Context, id int64) error {
	for i := range f.items {
		if f.items[i].ID == id {
			f.items = append(f.items[:i], f.items[i+1:]...)
			return nil
		}
	}
	return apperr.ErrNotFound
}

var _ repository.ApplicationRepository = (*fakeRepo)(nil)

func TestCreate_Validation(t *testing.T) {
	u := New(&fakeRepo{})

	if _, err := u.Create(context.Background(), models.Application{
		DocumentType: models.TypePassport,
	}); err != apperr.ErrValidation {
		t.Fatalf("expected ErrValidation for empty citizen_name, got %v", err)
	}

	if _, err := u.Create(context.Background(), models.Application{
		CitizenName:  "Иван",
		DocumentType: "wrong",
	}); err != apperr.ErrValidation {
		t.Fatalf("expected ErrValidation for bad document_type, got %v", err)
	}
}

func TestCreate_Get_UpdateStatus(t *testing.T) {
	fr := &fakeRepo{}
	u := New(fr)

	// create
	id, err := u.Create(context.Background(), models.Application{
		CitizenName:  "Иван",
		DocumentType: models.TypePassport,
		Data:         map[string]any{"series": "1234"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if id != 1 {
		t.Fatalf("want id=1, got %d", id)
	}

	// get
	got, err := u.GetByID(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	if got.CitizenName != "Иван" || got.DocumentType != models.TypePassport {
		t.Fatalf("unexpected entity: %+v", got)
	}

	// update status
	if err := u.UpdateStatus(context.Background(), id, models.StatusSubmitted); err != nil {
		t.Fatal(err)
	}
	got, _ = u.GetByID(context.Background(), id)
	if got.Status != models.StatusSubmitted {
		t.Fatalf("want status=submitted, got %s", got.Status)
	}
}
