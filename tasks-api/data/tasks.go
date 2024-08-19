package data

import (
	"encoding/csv"
	"errors"
	"github.com/hashicorp/go-hclog"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id" csv:"id"`
	Title       string    `json:"title" csv:"title"`
	Description string    `json:"description" csv:"description"`
	IsDone      bool      `json:"isDone" csv:"isDone"`
	AddedOn     time.Time `json:"-"`
	EditedOn    time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

type Tasks []*Task

type Handler struct {
	l         hclog.Logger
	TasksList Tasks
	csvH      *CsvHandler
}

func NewHandler(l hclog.Logger, csvH *CsvHandler) *Handler {
	var TasksList = Tasks{}
	return &Handler{l: l, TasksList: TasksList, csvH: csvH}
}

var TasksList1 = Tasks{}

func (d *Handler) ListTasks() Tasks {
	return d.csvH.ReadData()
}
func (d *Handler) ListSingleTask(id int) (*Task, error) {
	_, index, err := d.FindTaskById(id)
	if err != nil {
		return nil, err
	}
	return d.TasksList[index], nil
}

func (d *Handler) AddTask(t *Task) {
	t.Id = d.getNextId()
	t.AddedOn = time.Now()
	t.IsDone = false
	d.TasksList = append(d.TasksList, t)

	file, err := os.OpenFile("data/tasks.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write a new row to the CSV file
	var row = []string{strconv.Itoa(t.Id), t.Title, t.Description, "false\n"}

	err = writer.Write(row)
	if err != nil {
		panic(err)
	}
}

func (d *Handler) UpdateTask(id int, t *Task) error {
	_, i, err := d.FindTaskById(id)
	if err != nil {
		return err
	}
	t.Id = id
	d.TasksList[i] = t
	return nil
}

func (d *Handler) DeleteTask(id int) error {
	_, i, err := d.FindTaskById(id)
	if err != nil {
		return err
	}
	d.TasksList = append(d.TasksList[:i], d.TasksList[i+1:]...)
	return nil
}

func (d *Handler) getNextId() int {
	return d.TasksList[len(d.TasksList)-1].Id + 1
}

func (d *Handler) FindTaskById(id int) (*Task, int, error) {
	for i, t := range d.TasksList {
		if t.Id == id {
			return t, i, nil
		}
	}
	return &Task{}, -1, errors.New("unable to find task with this id")
}
