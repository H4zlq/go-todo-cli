package cmd

import (
	"fmt"

	"github.com/h4zlq/cli/model"
	db "github.com/h4zlq/cli/pkg/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress <task-id>",
	Short: "Mark a task as in progress",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a task id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		task := db.GetTask(id)

		if task.ID == 0 {
			fmt.Println("Task not found")
			return
		}

		db.UpdateTask(id, model.Task{
			ID:     task.ID,
			Status: "in-progress",
		})

		fmt.Printf("Task marked as in progress (ID: %d)\n", task.ID)
	},
}
