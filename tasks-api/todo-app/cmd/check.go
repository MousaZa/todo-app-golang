/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"net/http"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Marks the task with the specified id as done",

	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:9092/tasks/" + args[0]
		req, err := http.NewRequest(http.MethodPost, url, nil)
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
			c.Printf("Task with id: %v checked successfuly\n", args[0])
		} else {
			r := color.New(color.FgRed)
			r.Printf("Task with id: %v not found\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
