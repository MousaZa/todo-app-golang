package main

import (
	"github.com/MousaZa/todo-app/tasks-api/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

func main() {

	l := hclog.New(&hclog.LoggerOptions{Name: "tasks-api", Level: hclog.LevelFromString("DEBUG")})
	h := handlers.NewTaskHandler(l)

	sm := mux.NewRouter()

	sm.HandleFunc("/tasks", h.ListTasks).Methods(http.MethodGet)
	sm.HandleFunc("/tasks/{id:[0-9]+}", h.ListSingleTask).Methods(http.MethodGet)

	sm.HandleFunc("/tasks", h.AddTask).Methods(http.MethodPost)

	sm.HandleFunc("/tasks/{id:[0-9]+}", h.UpdateTask).Methods(http.MethodPut)

	err := http.ListenAndServe(":9091", sm)
	if err != nil {
		l.Error("Error starting Service", "error", err)
	}

}
