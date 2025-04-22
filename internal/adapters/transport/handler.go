package transport

import (
	"encoding/json"
	"net/http"
	"tasks/internal/ports/in"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service in.TaskService
}

func NewTaskHandler(s in.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.CreateTask(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := FromDomain(task)
	writeJSON(w, response)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := FromDomain(task)
	writeJSON(w, response)
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
