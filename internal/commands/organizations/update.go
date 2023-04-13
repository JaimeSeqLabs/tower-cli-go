package organizations

import (
	"fmt"

	"github.com/spf13/cobra"
)


var UpdateCmd = &cobra.Command{
	Use: "update",
	Short: "Update organization details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Updating orgs %v\n", args)
	},
}