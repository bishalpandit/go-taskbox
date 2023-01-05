package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bengadbois/flippytext"
	"github.com/bishalpandit/taskbox/cache"
	"github.com/bishalpandit/taskbox/constants"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
)



var view = new(constants.View)

func compareTime(t1 timestamppb.Timestamp, t2 timestamppb.Timestamp) bool {
	return (t1.GetSeconds() == t2.GetSeconds() && t1.GetNanos() > t2.GetNanos()) || 
	t1.GetSeconds() > t2.GetSeconds()
}

func compareOrder(order string, tasks *[]constants.Task, idx int) bool {
	if order == "asc" {
		return compareTime((*tasks)[idx].CreatedAt, (*tasks)[idx + 1].CreatedAt)
	} 
	return !compareTime((*tasks)[idx].CreatedAt, (*tasks)[idx + 1].CreatedAt)
}

func sortTasksByTimestamp(order string, tasks *[]constants.Task) {
	var temp constants.Task;

	for i := 0; i < len(*tasks) - 1; i++ {
		for j := 0; j < len(*tasks) - i - 1; j++ {
			if compareOrder(order, tasks, j) {
				temp = (*tasks)[j]
				(*tasks)[j] = (*tasks)[j + 1]
				(*tasks)[j + 1] = temp
			}
		}
	}
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks.",
	Long: `List all your tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		var allTasks []constants.Task
		var key string

		tasksCache, err := cache.GetTasks(key);
		if err != nil {
			b, err := os.ReadFile("./data/tasks.json")
			if err != nil {
				fmt.Println(err)
			}
	
			if err := json.Unmarshal(b, &allTasks)
			err != nil {
				fmt.Println(err)
			}
			
			if view.Sort != "" {
				if (view.Sort != "asc" || view.Sort != "desc") {
					err :=errors.New("invalid sort order. Use asc or desc")
					fmt.Println(err)
				}
			}
			cache.SetTasks(key, allTasks)
		} else {
			allTasks = tasksCache
		}

		data, err1 := json.MarshalIndent(allTasks, "", "   ")
		if err1 != nil {
			fmt.Println(err1)
		}
		
		f := *flippytext.New();
		f.TickerTime = time.Nanosecond/2;
		f.TickerCount = 4;

		f.Write(string(data[:]));
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&view.Sort, "sort", "s", "", "Sort based on creation time and priority");
}
