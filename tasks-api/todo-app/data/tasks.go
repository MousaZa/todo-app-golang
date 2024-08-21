package data

import (
	"fmt"
	"github.com/fatih/color"
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

type Tasks []Task

func (t *Task) PrintTask(d bool) {
	c := color.New(color.FgGreen)
	fmt.Printf("[")
	if t.IsDone {
		c.Printf("v")
	} else {
		fmt.Printf(" ")
	}
	fmt.Printf("] ")
	fmt.Printf("%v", t.Title)
	if d {
		fmt.Printf(" - %v", t.Description)
	}
	fmt.Printf("\n")
}
