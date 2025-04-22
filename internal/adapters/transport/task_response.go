package transport

import (
	"tasks/internal/domain"
	"time"
)

type TaskResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(task *domain.Task) TaskResponse {
	return TaskResponse{
		ID:        task.GetID(),
		Status:    string(task.GetStatus()),
		Result:    task.GetResult(),
		CreatedAt: task.GetCreatedAt(),
		UpdatedAt: task.GetUpdatedAt(),
	}
}

func (tr *TaskResponse) ToDomain() *domain.Task {
	return domain.NewTask(
		tr.ID,
		tr.Result,
		domain.TaskStatus(tr.Status),
		tr.CreatedAt,
		tr.UpdatedAt,
	)
}
