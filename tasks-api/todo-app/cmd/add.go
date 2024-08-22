/*
Copyright © 2024 Mousa Zeydan <mous.zeydan@gmail.com>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MousaZa/todo-app/tasks-api/todo-app/data"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"net/http"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the list",

	Run: func(cmd *cobra.Command, args []string) {
		t := data.Task{Title: args[0], Description: args[1]}
		j, err := json.Marshal(t)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		r := bytes.NewReader(j)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:9092/tasks", r)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		if res.StatusCode == http.StatusOK {
			c := color.New(color.FgCyan)
			c.Printf("Task was added successfuly\n")
		} else {
			r := color.New(color.FgRed)
			r.Printf("Error adding task, code: %v\n", res.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
