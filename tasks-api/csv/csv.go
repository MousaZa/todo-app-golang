package csv

import (
	"fmt"
	"github.com/MousaZa/todo-app/tasks-api/data"
	"github.com/gocarina/gocsv"
	"github.com/hashicorp/go-hclog"
	"os"
)

type Handler struct {
	l hclog.Logger
}

func NewHandler(l hclog.Logger) *Handler {
	return &Handler{l: l}
}

func (h *Handler) ReadData() data.Tasks {
	var TasksList = data.Tasks{}
	fmt.Println("sss")
	file, err := os.Open("csv/tasks.csv")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	err = gocsv.UnmarshalFile(file, &TasksList)
	if err != nil {
		fmt.Println(err)
	}
	return TasksList
}
