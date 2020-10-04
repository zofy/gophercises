package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is CLI task manager",
}

// Execute -
func Execute() error {
	return rootCmd.Execute()
}
