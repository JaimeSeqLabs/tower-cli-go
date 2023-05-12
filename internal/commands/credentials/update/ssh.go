package update_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateSSHCmd() *cobra.Command {

	sshCmd := &cobra.Command{
		Use:   "ssh",
		Short: "Update SSH workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewSSHCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(sshCmd)

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
