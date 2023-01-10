package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGetItem = &cobra.Command{
	Use:   "GetItem",
	Short: "GetItem from server",
	Long:  `It will receive message and related data from the DB`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println("GetItem")
	},
}

func init() {
	rootCmd.AddCommand(cmdGetItem)
}
