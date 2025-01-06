/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"task-tracker-cli/internal"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			fmt.Println("invalid task id")
		}

		err = internal.DeleteTaskCommand(int(id))

		if err != nil {
			fmt.Println(err)
		}
	},
}
