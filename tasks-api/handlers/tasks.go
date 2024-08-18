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
	l  hclog.Logger
	dh *data.Handler
}

func NewTaskHandler(l hclog.Logger, dh *data.Handler) *TaskHandler {
	return &TaskHandler{l: l, dh: dh}
}

func (h *TaskHandler) ListTasks(rw http.ResponseWriter, _ *http.Request) {
	t := h.dh.ListTasks()
	err := json.NewEncoder(rw).Encode(t)
	if err != nil {
		h.l.Error("Error Encoding Tasks", "error", err)
	}
}

func (h *TaskHandler) ListSingleTask(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	t, err := h.dh.ListSingleTask(id)
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

func (h *TaskHandler) AddTask(rw http.ResponseWriter, r *http.Request) {
	var t = data.Task{}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		h.l.Error("Error Decoding Tasks", "error", err)
		http.Error(rw, "Unable to Decode json", http.StatusBadRequest)
	}
	h.dh.AddTask(&t)
	rw.Write([]byte("Task Added Successfully\n"))
}

func (h *TaskHandler) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	var t = data.Task{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.l.Error("Error Getting Id", "error", err)
		http.Error(rw, "Wrong Id format", http.StatusBadRequest)
	}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		h.l.Error("Error Decoding Tasks", "error", err)
		http.Error(rw, "Unable to Decode json", http.StatusBadRequest)
	}
	err = h.dh.UpdateTask(id, &t)
	if err != nil {
		h.l.Error("Task with Id not found", "error", err)
		http.Error(rw, "Task with Id not found", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.l.Error("Error Getting Id", "error", err)
		http.Error(rw, "Wrong Id format", http.StatusBadRequest)
	}
	err = h.dh.DeleteTask(id)
	if err != nil {
		h.l.Error("Task with Id not found", "error", err)
		http.Error(rw, "Task with Id not found", http.StatusInternalServerError)
	}
}
