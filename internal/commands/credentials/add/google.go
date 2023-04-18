package add_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddGoogleCmd() *cobra.Command {

	googleCmd := &cobra.Command{
		Use:   "google",
		Short: "Add new Google workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewGoogleCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(googleCmd)

	googleCmd.Flags().StringP(
		"key", "k", "",
		"JSON file with the service account key",
	)
	cobra.MarkFlagRequired(googleCmd.Flags(), "key")

	return googleCmd
}
