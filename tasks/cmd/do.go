package cmd

import (
	"fmt"
	"strconv"
	"tasks/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var taskIds []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				taskIds = append(taskIds, id)
			}
		}
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Println("Error occured:", err)
			return

		}
		for _, id := range taskIds {
			if id < 0 || id > len(tasks) {
				fmt.Println("Inavalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error occured. %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
