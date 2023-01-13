package cmd

import (
	"client/messages"
	"client/structures"
	"strings"

	"github.com/spf13/cobra"
)

var cmdGetItem = &cobra.Command{
	Use:   "getItem [item to get]",
	Short: "Get Item from server",
	Long:  `It will print message on server side`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg := structures.MmgStructure{Message: "getItem", Body: strings.Join(args, " ")}
		messages.RabbitSend(msg)
	},
}

func init() {
	rootCmd.AddCommand(cmdGetItem)
}
