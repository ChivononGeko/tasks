package out

import (
	"context"
	"tasks/internal/domain"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	FindByID(ctx context.Context, id string) (*domain.Task, error)
}
