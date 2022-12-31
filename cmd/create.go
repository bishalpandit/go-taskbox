package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

type Task struct {
	id int
    title string
    description string
}

var p = new(Task)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long: `Create a new task`,
	Run: func(cmd *cobra.Command, args []string) {
		temp_id := rand.Int();
		p.id = temp_id;

		fmt.Println(*p);
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&p.title, "title", "t", "", "The title of a task");
	if err := CreateCmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err);
	}

	CreateCmd.Flags().StringVarP(&p.description, "description", "d", "", "The description of a task");
	if err := CreateCmd.MarkFlagRequired("description"); err != nil {
		fmt.Println(err);
	}
}
