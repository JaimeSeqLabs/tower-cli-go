package organizations

import (
	"fmt"

	"github.com/spf13/cobra"
)


var ViewCmd = &cobra.Command{
	Use: "view",
	Short: "Describe organization details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Viewing orgs %v\n", args)
	},
}