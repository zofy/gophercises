package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zofy/task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		for _, taskStr := range args {
			if _, err := db.CreateTask(taskStr); err != nil {
				fmt.Println("Could not create task: ", taskStr)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
