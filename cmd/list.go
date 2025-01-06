/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"task-tracker-cli/internal"

	"github.com/spf13/cobra"
)

var (
	sortByFlag                 string
	SORT_BY_STATUS_FLAG        = "status"
	SORT_BY_CREATION_TIME_FLAG = "creationTime"
	SORT_BY_UPDATE_TIME_FLAG   = "updateTime"
)

func isValidSortOption(option string, validOptions []string) bool {
	for _, validOption := range validOptions {
		if option == validOption {
			return true
		}
	}
	return false
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "Lists tasks",
	Long:      ``,
	ValidArgs: []string{string(internal.TASK_STATUS_TODO), string(internal.TASK_STATUS_IN_PROGRESS), string(internal.TASK_STATUS_DONE)},
	Args:      cobra.OnlyValidArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		validSortOptions := []string{SORT_BY_STATUS_FLAG, SORT_BY_CREATION_TIME_FLAG, SORT_BY_UPDATE_TIME_FLAG, ""}

		if !isValidSortOption(sortByFlag, validSortOptions) {
			return fmt.Errorf("invalid sort option: %s. Valid options are: %s", sortByFlag, strings.Join(validSortOptions, ", "))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var taskStatus internal.TaskStatus

		if len(args) == 1 {
			taskStatus = internal.TaskStatus(args[0])
		}

		var sortBy int
		var err error
		switch sortByFlag {
		case SORT_BY_STATUS_FLAG:
			sortBy = internal.SORT_BY_STATUS
		case SORT_BY_CREATION_TIME_FLAG:
			sortBy = internal.SORT_BY_CREATION_TIME
		case SORT_BY_UPDATE_TIME_FLAG:
			sortBy = internal.SORT_BY_UPDATE_TIME
		default:
			sortBy = 0
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		err = internal.ListTasksCommand(&taskStatus, sortBy)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	listCmd.Flags().StringVar(&sortByFlag, "sort", "", "Sort the items by 'status', 'creationTime', or 'updateTime'")
}
