package cmd

import (
	"client/messages"
	"client/structures"
	"strings"

	"github.com/spf13/cobra"
)

var cmdAddItem = &cobra.Command{
	Use:   "addItem [item to add]",
	Short: "Send new message to the server",
	Long:  `It will send new message to the server and save it to DB`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg := structures.MmgStructure{Message: "addItem", Body: strings.Join(args, " ")}
		messages.RabbitSend(msg)
	},
}

func init() {
	rootCmd.AddCommand(cmdAddItem)
}
