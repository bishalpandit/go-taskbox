package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Task struct {
	Id int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
	Tags []string `json:"tags"`
	CreatedAt timestamppb.Timestamp `json:"createdAt"`
	PriorityCode int16 `json:"priorityCode"`
	ExpectedTime int16 `json:"expectedTime"`
	DueTime int16 `json:"dueTime"`
}

var task = new(Task)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long: `Create a new task`,
	Run: func(cmd *cobra.Command, args []string) {
		task.Id = uuid.ClockSequence()
		task.CreatedAt = *timestamppb.Now()

		f, err2 := os.OpenFile("./data/tasks.json", os.O_RDWR, 0644)
		if err2 != nil {
			fmt.Println(err2)
		}

		b1, err3 := os.ReadFile("./data/tasks.json")
		if err3 != nil {
			fmt.Println(err3)
		}

		var allTasks []Task
		err := json.Unmarshal(b1, &allTasks)
		if err != nil {
			fmt.Println(err)
		}
		
		allTasks = append(allTasks, *task)

		b2, err1 := json.MarshalIndent(allTasks, "", "    ")
		if err1 != nil {
			fmt.Println(err1)
		}
		
		erry := os.WriteFile("./data/tasks.json", b2, 0644)
		if erry != nil {
			fmt.Println(err)
		}

		defer f.Close()
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&task.Title, "title", "t", "", "The title of a task");
	if err := CreateCmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err);
	}
	CreateCmd.Flags().StringVarP(&task.Description, "description", "d", "", "The description of a task");
	CreateCmd.Flags().StringSliceVar(&task.Tags, "tags", nil, "The tags related to a task");
	CreateCmd.Flags().Int16VarP(&task.PriorityCode, "priority", "p", -1, "The priority code " +
	"related to a task. 0 - High, 1 - Medium, 2 - Low");

	CreateCmd.Flags().Int16VarP(&task.ExpectedTime, "expectedTime", "e", -1, "The expected completion " +
	"time of a task.(in hours");
	CreateCmd.Flags().Int16VarP(&task.DueTime, "dueTime", "u", -1, "The due time duration " +
	"for a task(in hours).");
	// add validation for ET and DT
}
