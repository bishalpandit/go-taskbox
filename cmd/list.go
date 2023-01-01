package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bengadbois/flippytext"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks.",
	Long: `List all your tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("./data/tasks.json")
		if err != nil {
			fmt.Println(err)
		}

		var allTasks []Task
		if err := json.Unmarshal(b, &allTasks)
		err != nil {
			fmt.Println(err)
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

}
