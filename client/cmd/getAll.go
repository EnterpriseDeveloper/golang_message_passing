package cmd

import (
	"client/messages"
	"client/structures"

	"github.com/spf13/cobra"
)

var cmdGetAllItems = &cobra.Command{
	Use:   "getAllItems",
	Short: "Print item to the screen",
	Long:  `It will print all messages on server side`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, _ []string) {
		msg := structures.MmgStructure{Message: "getAllItems"}
		messages.RabbitSend(msg)
	},
}

func init() {
	rootCmd.AddCommand(cmdGetAllItems)
}
