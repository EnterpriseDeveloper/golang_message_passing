package cmd

import (
	"client/messages"
	"client/structures"

	"github.com/spf13/cobra"
)

var cmdGetItem = &cobra.Command{
	Use:   "getItem",
	Short: "Get Item from server",
	Long:  `It will print message on server side`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, _ []string) {
		msg := structures.MmgStructure{Message: "getItem"}
		messages.RabbitSend(msg)
	},
}

func init() {
	rootCmd.AddCommand(cmdGetItem)
}
