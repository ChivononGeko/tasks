package service

import (
	"context"
	"tasks/internal/domain"
	"tasks/internal/ports/in"
	"tasks/internal/ports/out"
	"time"

	"github.com/google/uuid"
)

type taskService struct {
	repo out.TaskRepository
}

func NewTaskService(repo out.TaskRepository) in.TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context) (*domain.Task, error) {
	id := uuid.New().String()
	now := time.Now()

	task := domain.NewTask(id, "", domain.StatusPending, now, now)

	err := s.repo.Save(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	return s.repo.FindByID(ctx, id)
}
