package data

import (
	"errors"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"isDone"`
	AddedOn     time.Time `json:"-"`
	EditedOn    time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

type Tasks []*Task

func ListTasks() Tasks {
	return TasksList
}
func ListSingleTask(id int) (*Task, error) {
	_, index, err := FindTaskById(id)
	if err != nil {
		return nil, err
	}
	return TasksList[index], nil
}

func AddTask(t *Task) {
	t.Id = getNextId()
	t.AddedOn = time.Now()
	t.IsDone = false
	TasksList = append(TasksList, t)
}

func UpdateTask(id int, t *Task) error {
	_, i, err := FindTaskById(id)
	if err != nil {
		return err
	}
	t.Id = id
	TasksList[i] = t
	return nil
}

func DeleteTask(id int) error {
	_, i, err := FindTaskById(id)
	if err != nil {
		return err
	}
	TasksList = append(TasksList[:i], TasksList[i+1:]...)
	return nil
}

func getNextId() int {
	return TasksList[len(TasksList)-1].Id + 1
}

func FindTaskById(id int) (*Task, int, error) {
	for i, t := range TasksList {
		if t.Id == id {
			return t, i, nil
		}
	}
	return &Task{}, -1, errors.New("unable to find task with this id")
}

var TasksList = Tasks{
	&Task{
		Id:          1,
		Title:       "Buy milk",
		Description: "Buy milk from the market",
		IsDone:      false,
		AddedOn:     time.Now(),
		EditedOn:    time.Now(),
	},
	&Task{
		Id:          2,
		Title:       "Finish the project",
		Description: "Do the last changes and push the code",
		IsDone:      false,
		AddedOn:     time.Now(),
		EditedOn:    time.Now(),
	},
}
