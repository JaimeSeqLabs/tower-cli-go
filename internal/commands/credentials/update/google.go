package update_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateGoogleCmd() *cobra.Command {

	googleCmd := &cobra.Command{
		Use:   "google",
		Short: "Update Google workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewGoogleCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(googleCmd)

	googleCmd.Flags().StringP(
		"key", "k", "",
		"JSON file with the service account key",
	)
	cobra.MarkFlagRequired(googleCmd.Flags(), "key")

	return googleCmd
}
