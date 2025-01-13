package cmd

import (
	"fmt"
	"strings"

	"github.com/h4zlq/cli/model"
	db "github.com/h4zlq/cli/pkg/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update <id> <task>",
	Short: "Update a task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("requires an id and a task description")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		task := model.Task{Description: strings.Join(args[1:], " ")}

		db.UpdateTask(id, task)

		fmt.Printf("Task updated successfully (ID: %s)\n", id)
	},
}
