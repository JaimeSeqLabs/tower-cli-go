package organizations

import "github.com/spf13/cobra"

func NewOrganizationsCmd() *cobra.Command {

	orgCmd := &cobra.Command{
		Use:   "organizations",
		Short: "Manage organizations",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	orgCmd.AddCommand(
		NewListCmd(),
		NewDeleteCmd(),
		NewAddCmd(),
		NewUpdateCmd(),
		NewViewCmd(),
	)

	return orgCmd
}
