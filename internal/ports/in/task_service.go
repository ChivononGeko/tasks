package in

import (
	"context"
	"tasks/internal/domain"
)

type TaskService interface {
	CreateTask(ctx context.Context) (*domain.Task, error)
	GetTask(ctx context.Context, id string) (*domain.Task, error)
}
