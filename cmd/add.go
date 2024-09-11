package cmd

import (
	"fmt"

	"github.com/lorenzomene/go-todo-cli/todo"
	"github.com/spf13/cobra"
)

var todos = todo.Todos{}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "add a task to the todo",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		success, err := todos.AddTask(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		if success {
			fmt.Println("Task added:", task)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
