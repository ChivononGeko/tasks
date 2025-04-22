package repository

import (
	"tasks/internal/domain"
	"time"
)

type taskRecord struct {
	ID        string
	Status    string
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(task *domain.Task) *taskRecord {
	return &taskRecord{
		ID:        task.GetID(),
		Status:    string(task.GetStatus()),
		Result:    task.GetResult(),
		CreatedAt: task.GetCreatedAt(),
		UpdatedAt: task.GetUpdatedAt(),
	}
}

func (r *taskRecord) toDomain() *domain.Task {
	return domain.NewTask(
		r.ID,
		r.Result,
		domain.TaskStatus(r.Status),
		r.CreatedAt,
		r.UpdatedAt,
	)
}
