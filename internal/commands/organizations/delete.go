package organizations

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an organization",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting orgs %v\n", args)
	},
}
