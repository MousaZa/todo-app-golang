package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/hashicorp/go-hclog"
	"log"
	"os"
	"strconv"
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

func (h *CsvHandler) AddTask(t *Task) error {
	file, err := os.OpenFile("data/csv/tasks.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		h.l.Error("error1", err)
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write a new row to the CSV file
	var row = []string{strconv.Itoa(t.Id), t.Title, t.Description, "false"}

	err = writer.Write(row)
	if err != nil {
		return err
	}
	return nil
}

func (h *CsvHandler) UpdateTask(t *Task, id string) error {
	file, err := os.OpenFile("data/csv/tasks.csv", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		h.l.Error("error", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		h.l.Info("Closing file")
		if err != nil {

		}
	}(file)
	//Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write a new row to the CSV file
	var row = []string{strconv.Itoa(t.Id), t.Title, t.Description, "false"}

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()

	if err != nil {
		h.l.Error("error reading file", "error", err)
		return err
	}
	index := -1
	for i, _ := range rows {
		if rows[i][0] == id {
			index = i
		}
	}
	if index == -1 {
		return errors.New("task with id not found")
	}
	rows[index] = row
	if err := os.Truncate("data/csv/tasks.csv", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	err = writer.WriteAll(rows)
	if err != nil {
		h.l.Error("error", err)
		return err
	}
	return nil
}

func (h *CsvHandler) DeleteTask(id string) error {
	file, err := os.OpenFile("data/csv/tasks.csv", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		h.l.Error("error", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		h.l.Info("Closing file")
		if err != nil {

		}
	}(file)
	//Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write a new row to the CSV file

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()

	if err != nil {
		h.l.Error("error reading file", "error", err)
		return err
	}
	index := -1
	for i, _ := range rows {
		if rows[i][0] == id {
			index = i
		}
	}
	h.l.Debug("index", index)
	if index == -1 {
		return errors.New("task with id not found")
	}
	h.l.Debug("index", index)
	rows = append(rows[:index], rows[index+1:]...)
	if err := os.Truncate("data/csv/tasks.csv", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	err = writer.WriteAll(rows)
	if err != nil {
		h.l.Error("error", err)
		return err
	}
	return nil
}
