package add_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddSSHCmd() *cobra.Command {

	sshCmd := &cobra.Command{
		Use:   "ssh",
		Short: "Add new SSH workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewSSHCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(sshCmd)

	sshCmd.Flags().StringP(
		"key", "k", "",
		"SSH private key file",
	)
	cobra.MarkFlagRequired(sshCmd.Flags(), "key")

	sshCmd.Flags().StringP(
		"passphrase", "p", "",
		"Passphrase associated with the private key",
	)

	return sshCmd
}
