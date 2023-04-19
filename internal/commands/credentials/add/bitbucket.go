package add_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddBitBucketCmd() *cobra.Command {

	bitBuckCmd := &cobra.Command{
		Use:   "bitbucket",
		Short: "Add new Bitbucket workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewBitBucketCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(bitBuckCmd)

	// TODO: shorthand -u collides with global -u url
	bitBuckCmd.Flags().String(
		"username", "",
		"Bitbucket username",
	)
	cobra.MarkFlagRequired(bitBuckCmd.Flags(), "username")

	bitBuckCmd.Flags().StringP(
		"password", "p", "",
		"Bitbucket App password",
	)
	cobra.MarkFlagRequired(bitBuckCmd.Flags(), "password")

	bitBuckCmd.Flags().String(
		"base-url", "",
		"Repository base URL",
	)

	return bitBuckCmd
}
