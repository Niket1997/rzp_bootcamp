package cmd

import "github.com/spf13/cobra"

// RootCmd variable to export basic cli
var RootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "tasks is a CLI task manager",
}
