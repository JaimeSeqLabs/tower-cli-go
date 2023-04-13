package organizations

import (
	"fmt"

	"github.com/spf13/cobra"
)


var ListCmd = &cobra.Command{
	Use: "list",
	Short: "List available organizations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing orgs %v\n", args)
	},
}