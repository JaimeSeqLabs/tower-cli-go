package update_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateBitBucketCmd() *cobra.Command {

	bitBuckCmd := &cobra.Command{
		Use:   "bitbucket",
		Short: "Update Bitbucket workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewBitBucketCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(bitBuckCmd)

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
