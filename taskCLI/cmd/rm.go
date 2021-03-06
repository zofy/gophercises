package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zofy/task/db"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "deletes task from your list",
	Run: func(cmd *cobra.Command, args []string) {
		var idxs []int
		for _, arg := range args {
			idx, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse: %s\n", arg)
			} else {
				idxs = append(idxs, idx)
			}
		}
		if err := db.DeleteTasks(idxs); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
