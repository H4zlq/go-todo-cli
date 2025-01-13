package cmd

import (
	"fmt"

	db "github.com/h4zlq/cli/pkg/database"
	utils "github.com/h4zlq/cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:       "list <task-status>",
	Short:     "List all tasks or tasks with a specific status",
	ValidArgs: []string{"todo", "in-progress", "done"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			status := args[0]
			if status != "todo" && status != "in-progress" && status != "done" {
				return fmt.Errorf("invalid status. Must be one of: todo, in-progress, done")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.GetTasks()

		if len(tasks) == 0 {
			fmt.Println("You don't have any tasks yet")
			return
		}

		if len(args) > 0 {
			status := args[0]
			tasks := db.GetTasksByStatus(status)

			if len(tasks) == 0 {
				fmt.Println("No tasks found with status", status)
				return
			}

			fmt.Println("You have the following tasks with status:", status)

			for _, task := range tasks {
				fmt.Printf("%d. %s (%s)\n", task.ID, task.Description, utils.FormatStatus(task.Status))
			}

			return
		}

		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s (%s)\n", task.ID, task.Description, utils.FormatStatus(task.Status))
		}
	},
}
