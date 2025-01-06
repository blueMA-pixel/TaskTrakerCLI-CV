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

const invalidTaskIdErrorMessage = "invalid task id"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update the task description",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 32)

		if err != nil {
			fmt.Println(invalidTaskIdErrorMessage)
		}

		newDescription := args[1]

		err = internal.UpdateTaskCommand(int(id), newDescription)

		if err != nil {
			fmt.Println(err)
		}
	},
}

var markInProgressCommand = &cobra.Command{
	Use:   "marke-in-progress",
	Short: "mark a task in progress",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 32)

		if err != nil {
			fmt.Println(invalidTaskIdErrorMessage)
		}

		err = internal.MarkCommand(int(id), internal.TASK_STATUS_IN_PROGRESS)

		if err != nil {
			fmt.Println(err)
		}

	},
}

var markDoneCommand = &cobra.Command{
	Use:   "marke-done",
	Short: "mark a task as done",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 32)

		if err != nil {
			fmt.Println(invalidTaskIdErrorMessage)
		}

		err = internal.MarkCommand(int(id), internal.TASK_STATUS_DONE)

		if err != nil {
			fmt.Println(err)
		}

	},
}
