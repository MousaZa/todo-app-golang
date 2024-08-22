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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
