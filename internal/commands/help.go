package commands

import "github.com/spf13/cobra"

var helpCmd = &cobra.Command {
	Use:   "help",
	Short: "Show cli help",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}

