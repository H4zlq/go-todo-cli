package cmd

import (
	"fmt"
	"strings"

	"github.com/h4zlq/cli/model"
	db "github.com/h4zlq/cli/pkg/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add <task-description>",
	Short: "Add a new task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a task description")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		task := model.Task{Description: strings.Join(args, " "), Status: "todo"}

		db.AddTask(task)

		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	},
}
