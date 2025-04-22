package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	TaskHandler *TaskHandler
	chiRouter   chi.Router
}

func NewRouter(taskHandler *TaskHandler) *Router {
	r := chi.NewRouter()

	r.Post("/tasks", taskHandler.CreateTask)
	r.Get("/tasks/{id}", taskHandler.GetTask)

	return &Router{
		TaskHandler: taskHandler,
		chiRouter:   r,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chiRouter.ServeHTTP(w, req)
}
