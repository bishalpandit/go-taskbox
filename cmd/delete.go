package cmd

import (
	"fmt"
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/bishalpandit/taskbox/constants"
)

var id int

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by id",
	Long: `Delete a task by id`,
	Run: func(cmd *cobra.Command, args []string) {
		var allTasks []constants.Task
		var newTasks []constants.Task

		b, err := os.ReadFile("./data/tasks.json")
		if err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal(b, &allTasks)
		err != nil {
			fmt.Println(err)
		}

		if id == -1 {
			err := errors.New("Task id missing")
			fmt.Println(err)
			os.Exit(5)
		}

		for _, task := range allTasks {
			if task.Id != id {
				newTasks = append(newTasks, task)
			}
		}

		b1, err1 := json.MarshalIndent(newTasks, "", "    ")
		if err1 != nil {
			fmt.Println(err1)
		}
		
		err2 := os.WriteFile("./data/tasks.json", b1, 0644)
		if err2 != nil {
			fmt.Println(err2)
		}

	},
}

func init() {
	DeleteCmd.Flags().IntVarP(&id, "id", "i", -1, "ID of task to be deleted");
}
