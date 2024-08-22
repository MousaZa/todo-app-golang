package data

import (
	"errors"
	"github.com/hashicorp/go-hclog"
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
	l    hclog.Logger
	csvH *CsvHandler
}

func NewHandler(l hclog.Logger, csvH *CsvHandler) *Handler {
	return &Handler{l: l, csvH: csvH}
}

func (d *Handler) ListTasks() Tasks {
	return d.csvH.ReadData()
}
func (d *Handler) ListSingleTask(id int) (*Task, error) {
	t := d.csvH.ReadData()
	_, index, err := d.FindTaskById(t, id)
	if err != nil {
		return nil, err
	}
	return t[index], nil
}

func (d *Handler) AddTask(t *Task) {
	tasks := d.csvH.ReadData()
	t.Id = d.getNextId(tasks)
	t.AddedOn = time.Now()
	t.IsDone = false
	err := d.csvH.AddTask(t)
	if err != nil {
		d.l.Error("Error Adding Task", err)
	}

}

func (d *Handler) UpdateTask(id string, t *Task) error {
	//tasks := d.csvH.ReadData()
	//_, i, err := d.FindTaskById(tasks, id)
	//d.l.Debug("i", i)
	//if err != nil {
	//	return err
	//}
	t.Id, _ = strconv.Atoi(id)
	err := d.csvH.UpdateTask(t, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Handler) DeleteTask(id string) error {
	err := d.csvH.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
func (d *Handler) CheckTask(id string) error {
	err := d.csvH.CheckTask(id)
	if err != nil {
		return err
	}
	return nil
}
func (d *Handler) getNextId(t Tasks) int {
	return t[len(t)-1].Id + 1
}

func (d *Handler) FindTaskById(t Tasks, id int) (*Task, int, error) {
	for i, ta := range t {
		if ta.Id == id {
			d.l.Debug("id", id, "ta.Id", ta.Id)
			return ta, i, nil
		}
	}
	return &Task{}, -1, errors.New("unable to find task with this id")
}
