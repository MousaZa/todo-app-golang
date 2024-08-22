/*
Copyright © 2024 Mousa Zeydan <mous.zeydan@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"net/http"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task with specific id",
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:9092/tasks/" + args[0]
		req, err := http.NewRequest(http.MethodDelete, url, nil)
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
			c.Printf("Task with id: %v deleted successfuly\n", args[0])
		} else {
			r := color.New(color.FgRed)
			r.Printf("Task with id: %v not found\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
