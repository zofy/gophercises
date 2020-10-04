package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zofy/task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ToDoTasks()
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to finsih!")
			return
		}
		for i, t := range tasks {
			fmt.Printf("%d: %s\n", i+1, t.Value)
		}
	},
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "list completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.DoneTasks()
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		if len(tasks) == 0 {
			fmt.Println("You haven't finished any task today.")
			return
		}
		fmt.Println("You've completed following tasks today:")
		for i, t := range tasks {
			fmt.Printf("%d: %s\n", i+1, t.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completedCmd)
}
