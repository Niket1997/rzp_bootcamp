package cmd

import (
	"fmt"
	"os"
	"strings"
	"tasks/db"

	"github.com/spf13/cobra"
)

// addCmd to add a task
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add is used to add task to tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your tasks list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
