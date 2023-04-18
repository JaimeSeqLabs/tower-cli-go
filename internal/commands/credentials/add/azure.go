package add_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddAzureCmd() *cobra.Command {

	azureCmd := &cobra.Command{
		Use:   "azure",
		Short: "Add new Azure workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewAzureCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(azureCmd)

	azureCmd.Flags().String(
		"batch-key", "",
		"Azure batch account key",
	)
	cobra.MarkFlagRequired(azureCmd.Flags(), "batch-key")

	azureCmd.Flags().String(
		"batch-name", "",
		"Azure batch account name",
	)
	cobra.MarkFlagRequired(azureCmd.Flags(), "batch-name")

	azureCmd.Flags().String(
		"storage-key", "",
		"Azure blob storage account key",
	)
	cobra.MarkFlagRequired(azureCmd.Flags(), "storage-key")

	azureCmd.Flags().String(
		"storage-name", "",
		"Azure blob storage account name",
	)
	cobra.MarkFlagRequired(azureCmd.Flags(), "storage-name")

	return azureCmd
}
