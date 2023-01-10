package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdRemoveItem = &cobra.Command{
	Use:   "RemoveItem [item to remove]",
	Short: "Print anything to the screen",
	Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RemoveItem: " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(cmdRemoveItem)
}
