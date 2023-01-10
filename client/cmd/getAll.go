package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGetAllItems = &cobra.Command{
	Use:   "GetAllItems [string to print]",
	Short: "Print anything to the screen",
	Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println("GetAllItems")
	},
}

func init() {
	rootCmd.AddCommand(cmdGetAllItems)
}
