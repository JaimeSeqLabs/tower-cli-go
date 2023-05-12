package update_cmd

import (
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateAzureCmd() *cobra.Command {

	azureCmd := &cobra.Command{
		Use:   "azure",
		Short: "Update Azure workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewAzureCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(azureCmd)
	common_flags.AddAzureKeysMandatoryFlags(azureCmd)

	return azureCmd
}
