package main

import (
	"context"
	"github.com/MousaZa/todo-app/tasks-api/data"
	"github.com/MousaZa/todo-app/tasks-api/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := hclog.New(&hclog.LoggerOptions{Name: "tasks-api", Level: hclog.LevelFromString("DEBUG")})
	csvH := data.NewCsvHandler(l)
	dh := data.NewHandler(l, csvH)
	h := handlers.NewTaskHandler(l, dh)

	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/tasks", h.ListTasks)
	getR.HandleFunc("/tasks/{id:[0-9]+}", h.ListSingleTask)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/tasks", h.AddTask)
	postR.HandleFunc("/tasks/{id:[0-9]+}", h.CheckTask)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/tasks/{id:[0-9]+}", h.UpdateTask)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/tasks/{id:[0-9]+}", h.DeleteTask)

	s := http.Server{
		Addr:         ":9092",
		Handler:      sm,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Error("error starting server", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Info("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
