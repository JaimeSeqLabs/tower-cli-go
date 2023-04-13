package organizations

import "github.com/spf13/cobra"

var OrganizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Manage organizations",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	OrganizationsCmd.AddCommand(
		ListCmd,
		DeleteCmd,
		AddCmd,
		UpdateCmd,
		ViewCmd,
	)
}
