/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/MousaZa/todo-app/tasks-api/todo-app/data"
	"github.com/fatih/color"
	"github.com/tidwall/pretty"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// listSingleCmd represents the listSingle command
var listSingleCmd = &cobra.Command{
	Use:   "listSingle",
	Short: "Lists a single task with a specific id",
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:9092/tasks/" + args[0]
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		//body, err := io.ReadAll(res.Body)
		if res.StatusCode != http.StatusOK {
			r := color.New(color.FgRed)
			r.Printf("Task with id: %v not found\n", args[0])
			return
		}
		if Debug {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			tasks := pretty.Pretty(body)
			fmt.Printf("%s", tasks)
		} else {
			t := &data.Task{}
			json.NewDecoder(res.Body).Decode(t)

			t.PrintTask(Verbose)
		}
		/*result := pretty.Pretty(body)
		fmt.Printf("Task :%s\n", result)*/

	},
}

func init() {
	rootCmd.AddCommand(listSingleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listSingleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listSingleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
