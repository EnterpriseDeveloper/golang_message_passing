package cmd

import (
	"client/messages"
	"client/structures"
	"strings"

	"github.com/spf13/cobra"
)

var cmdRemoveItem = &cobra.Command{
	Use:   "removeItem [item to remove]",
	Short: "Remove Item from DB",
	Long:  `It will remove item from server DB`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg := structures.MmgStructure{Message: "removeItem", Body: strings.Join(args, " ")}
		messages.RabbitSend(msg)
	},
}

func init() {
	rootCmd.AddCommand(cmdRemoveItem)
}
