package organizations

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new organization",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding orgs %v\n", args)
	},
}
