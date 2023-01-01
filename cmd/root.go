package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskbox",
	Short: "Task management CLI-based utility mainly aimed for terminal lovers with auto-priority scheduling.",
	Long: `Task management CLI-based utility mainly aimed for terminal lovers with auto-priority scheduling.`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(CreateCmd)
	rootCmd.AddCommand(ListCmd)
}


