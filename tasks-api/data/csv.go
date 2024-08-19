package data

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/hashicorp/go-hclog"
	"os"
)

type CsvHandler struct {
	l hclog.Logger
}

func NewCsvHandler(l hclog.Logger) *CsvHandler {
	return &CsvHandler{l: l}
}

func (h *CsvHandler) ReadData() Tasks {
	var TasksList = Tasks{}
	fmt.Println("sss")
	file, err := os.Open("data/csv/tasks.csv")
	if err != nil {
		panic(err)
	}

	err = gocsv.UnmarshalFile(file, &TasksList)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		fmt.Println(err)
	}
	return TasksList
}
