package main

import (
	"context"
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
	h := handlers.NewTaskHandler(l)

	sm := mux.NewRouter()

	sm.HandleFunc("/tasks", h.ListTasks).Methods(http.MethodGet)
	sm.HandleFunc("/tasks/{id:[0-9]+}", h.ListSingleTask).Methods(http.MethodGet)

	sm.HandleFunc("/tasks", h.AddTask).Methods(http.MethodPost)

	sm.HandleFunc("/tasks/{id:[0-9]+}", h.UpdateTask).Methods(http.MethodPut)

	s := http.Server{
		Addr:         ":9090",
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
