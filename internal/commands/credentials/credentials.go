package credentials

import "github.com/spf13/cobra"


func NewCredentialsCmd() *cobra.Command {
	
	credsCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Manage credentials",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	credsCmd.AddCommand(
		NewListCmd(),
	)

	return credsCmd
}