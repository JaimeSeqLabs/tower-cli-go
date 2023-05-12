package secrets

import "github.com/spf13/cobra"

func NewSecretsCmd() *cobra.Command {

	secretsCmd := &cobra.Command{
		Use:   "secrets",
		Short: "Manage workspace secrets",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	secretsCmd.AddCommand(
		NewAddCmd(),
		NewListCmd(),
		NewViewCmd(),
		NewUpdateCmd(),
		NewDeleteCmd(),
	)

	return secretsCmd
}
