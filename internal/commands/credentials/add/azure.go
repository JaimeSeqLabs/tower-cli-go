package add_cmd

import (
	"tower-cli-go/internal/commands/common_flags"
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
	common_flags.AddAzureKeysMandatoryFlags(azureCmd)

	return azureCmd
}
