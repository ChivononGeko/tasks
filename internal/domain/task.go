package domain

import "time"

type TaskStatus string

const (
	StatusPending   TaskStatus = "PENDING"
	StatusRunning   TaskStatus = "RUNNING"
	StatusCompleted TaskStatus = "COMPLETED"
	StatusFailed    TaskStatus = "FAILED"
)

type Task struct {
	id        string
	status    TaskStatus
	result    string
	createdAt time.Time
	updatedAt time.Time
}

func NewTask(id, result string, status TaskStatus, createdAt, updatedAt time.Time) *Task {
	return &Task{
		id:        id,
		status:    status,
		result:    result,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (t *Task) GetID() string           { return t.id }
func (t *Task) GetStatus() TaskStatus   { return t.status }
func (t *Task) GetResult() string       { return t.result }
func (t *Task) GetCreatedAt() time.Time { return t.createdAt }
func (t *Task) GetUpdatedAt() time.Time { return t.updatedAt }
