/*
Copyright © 2024 Mousa Zeydan <mous.zeydan@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/MousaZa/todo-app/tasks-api/todo-app/data"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9092/tasks", nil)
		if err != nil {
			fmt.Printf("err :%v\n", err)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("err :%v\n", err)
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
			tasks := data.Tasks{}
			err = json.NewDecoder(res.Body).Decode(&tasks)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, t := range tasks {
				t.PrintTask(Verbose)
			}
		}
	},
}

var Verbose bool
var Debug bool

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}
