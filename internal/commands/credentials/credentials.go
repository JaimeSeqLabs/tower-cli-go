package credentials

import (
	add_cmd "tower-cli-go/internal/commands/credentials/add"
	update_cmd "tower-cli-go/internal/commands/credentials/update"

	"github.com/spf13/cobra"
)

func NewCredentialsCmd() *cobra.Command {

	credsCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Manage credentials",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	credsCmd.AddCommand(
		NewListCmd(),
		add_cmd.NewAddCmd(),
		update_cmd.NewUpdateCmd(),
	)

	return credsCmd
}
