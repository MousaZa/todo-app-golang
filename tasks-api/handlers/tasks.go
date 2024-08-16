package handlers

import (
	"encoding/json"
	"github.com/MousaZa/todo-app/tasks-api/data"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	l hclog.Logger
}

func NewTaskHandler(l hclog.Logger) *TaskHandler {
	return &TaskHandler{l: l}
}

func (h *TaskHandler) ListTasks(rw http.ResponseWriter, r *http.Request) {
	t := data.ListTasks()
	err := json.NewEncoder(rw).Encode(t)
	if err != nil {
		h.l.Error("Error Encoding Tasks", "error", err)
	}
}

func (h *TaskHandler) ListSingleTask(rw http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	t, err := data.ListSingleTask(id)
	if err != nil {
		h.l.Error("Error Finding task", "id", id, "error", err)
		http.Error(rw, "Error Finding task", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(rw).Encode(t)
	if err != nil {
		h.l.Error("Error Encoding Tasks", "error", err)
		http.Error(rw, "Error Encoding Tasks", http.StatusInternalServerError)
		return
	}
}
