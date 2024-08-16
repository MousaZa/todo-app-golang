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
	index, err := FindTaskById(id)
	if err != nil {
		return nil, err
	}
	return TasksList[index], nil
}
func FindTaskById(id int) (int, error) {
	for i, t := range TasksList {
		if t.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("unable to find task with this id")
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
