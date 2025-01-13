package cmd

import (
	"fmt"

	"github.com/h4zlq/cli/model"
	db "github.com/h4zlq/cli/pkg/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete <task-id>",
	Short: "Delete a task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a task id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		task := model.Task{}

		db.DeleteTask(id, task)

		fmt.Printf("Task deleted successfully (ID: %s)\n", id)
	},
}
