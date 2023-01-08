package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
	"strings"

	"github.com/bengadbois/flippytext"
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

func takeNTasks(n int32, tasks *[]constants.Task) []constants.Task {
	if n > int32(len(*tasks)) {
		return *tasks
	}

	return (*tasks)[:n]
}

func parseTags(tagsStr string) ([]string, error) {
	return strings.Split(tagsStr, " "), nil
}

func filterByTags(tasks *[]constants.Task, tags []string) []constants.Task {
	tagsCount := 0
	var result []constants.Task

	set := make(map[string]struct{})
	for _, s := range tags {
		strTrimmed := strings.TrimSpace(s)
		set[strTrimmed] = struct{}{}
	}

	for _, task := range *tasks {
		tagsPresent := task.Tags

		for _, tag := range tagsPresent {
			tagTrimmed := strings.TrimSpace(tag)
			if _, ok := set[tagTrimmed]; ok {
				tagsCount++
			}
		}

		if tagsCount == len(set) {
			result = append(result, task)
		}
		tagsCount = 0
	}

	return result
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks.",
	Long: `List all your tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		var allTasks []constants.Task

		b, err := os.ReadFile("./data/tasks.json")
		if err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal(b, &allTasks)
		err != nil {
			fmt.Println(err)
		}
		
		if view.Sort != "" {
			if view.Sort != "asc" && view.Sort != "desc" {
				err :=errors.New("invalid sort order. Use asc or desc")
				fmt.Println(err)
				os.Exit(1)
			}

			sortTasksByTimestamp(view.Sort, &allTasks);
		}

		if view.Tail != -1 {
			if view.Tail < 1 {
				err :=errors.New("invalid input. Give a value > 0")
				fmt.Println(err)
				os.Exit(2)
			}

			allTasks = takeNTasks(view.Tail, &allTasks)
		}

		if view.TagFilter != "" {
			tags, err := parseTags(view.TagFilter)
			if err != nil {
				err := errors.New("Tags parsing error!")
				fmt.Println(err)
				os.Exit(3)
			}

			allTasks = filterByTags(&allTasks, tags)
		}

		data, err1 := json.MarshalIndent(&allTasks, "", "   ")
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
	ListCmd.Flags().StringVarP(&view.Sort, "sort", "s", "", "Sort based on creation time and priority");
	ListCmd.Flags().StringVarP(&view.TagFilter, "include", "i", "", "Filter tasks by tags assigned to them");
	ListCmd.Flags().Int32VarP(&view.Tail, "tail", "t", -1, "Retrieve last N tasks");
}
