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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the task with the specified id",
	Run: func(cmd *cobra.Command, args []string) {
		t := data.Task{Title: args[1], Description: args[2]}
		j, err := json.Marshal(t)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		r := bytes.NewReader(j)
		url := "http://localhost:9092/tasks/" + args[0]
		req, err := http.NewRequest(http.MethodPut, url, r)
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
			c.Printf("Task was updated successfuly\n")
		} else {
			r := color.New(color.FgRed)
			r.Printf("Error updating task, code: %v\n", res.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
