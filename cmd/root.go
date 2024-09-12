package cmd

import (
	"fmt"
	"os"

	"github.com/lorenzomene/go-todo-cli/todo"
	"github.com/spf13/cobra"
)

var todos = &todo.Todos{}
var csvFile = "tasks.csv"

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo list CLI app",
	Long:  `simple cli app with go`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := todos.LoadFromCSV(csvFile)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
}
