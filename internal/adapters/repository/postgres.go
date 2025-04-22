package repository

import (
	"context"
	"database/sql"
	"tasks/internal/domain"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Save(ctx context.Context, task *domain.Task) error {
	rec := fromDomain(task)

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO tasks (id, status, result, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`, rec.ID, rec.Status, rec.Result, rec.CreatedAt, rec.UpdatedAt)

	return err
}

func (r *PostgresRepo) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, status, result, created_at, updated_at
		FROM tasks WHERE id=$1
	`, id)

	var rec taskRecord
	err := row.Scan(&rec.ID, &rec.Status, &rec.Result, &rec.CreatedAt, &rec.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return rec.toDomain(), nil
}
