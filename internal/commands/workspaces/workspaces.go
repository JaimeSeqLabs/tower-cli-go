package workspaces

import "github.com/spf13/cobra"

func NewWorkspacesCmd() *cobra.Command {

	wspCmd := &cobra.Command{
		Use:   "workspaces",
		Short: "Manage workspaces",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	wspCmd.AddCommand(
		NewListWorkspacesCmd(),
		NewDeleteWorkspacesCmd(),
		NewAddWorkspacesCmd(),
		NewUpdateWorkspacesCmd(),
		NewViewWorkspacesCmd(),
		NewLeaveWorkspacesCmd(),
	)

	return wspCmd
}
