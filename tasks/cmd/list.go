package cmd

import (
	"fmt"
	"os"
	"tasks/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your created tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("error occured:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You don't have any tasks")
			return
		}
		fmt.Println("Your tasks are:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value )
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
