package common_flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddOptionalWorkspaceFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(
		"workspace", "w", "",
		"Workspace numeric identifier (TOWER_WORKSPACE_ID as default) or workspace reference as \"OrganizationName/WorkspaceName\"",
	)
	viper.BindEnv("workspace", "TOWER_WORKSPACE_ID")
	viper.BindPFlag("workspace", cmd.Flags().Lookup("workspace"))
}

func AddOrganizationDataFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(
		"description", "d", "",
		"Organization description",
	)
	cmd.Flags().StringP(
		"location", "l", "",
		"Organization location",
	)
	cmd.Flags().StringP(
		"website", "w", "",
		"Organization website URL",
	)
}

func AddOrganizationRefFlag(cmd *cobra.Command) {
	cmd.Flags().Int64P(
		"id", "i", -1,
		"Organization unique ID",
	)
	cmd.Flags().StringP(
		"name", "n", "",
		"Organization name",
	)
	cmd.MarkFlagsMutuallyExclusive("id", "name")
}

func AddWorkspaceRefFlag(cmd *cobra.Command) {
	cmd.Flags().Int64P(
		"id", "i", -1,
		"Workspace unique ID",
	)
	cmd.Flags().StringP(
		"name", "n", "",
		"Workspace name in the form of \"organizationName/workspaceName\"",
	)
	cmd.MarkFlagsMutuallyExclusive("id", "name")
}

func AddSecretRefFlags(cmd *cobra.Command) {
	cmd.Flags().Int64P(
		"id", "i", -1,
		"Secret unique ID",
	)
	cmd.Flags().StringP(
		"name", "n", "",
		"Secret name",
	)
}

func AddAwsKeysFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(
		"access-key", "a", "",
		"The AWS access key required to access the desired service",
	)
	cmd.Flags().StringP(
		"secret-key", "s", "",
		"The AWS secret key required to access the desired service",
	)
}

func AddOptionalCredentialsRefFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("name", "n", "", "Credentials name")
	cmd.Flags().StringP("id", "i", "", "Credentials unique id")
}

func AddAzureKeysMandatoryFlags(cmd *cobra.Command) {
	cmd.Flags().String(
		"batch-key", "",
		"Azure batch account key",
	)
	cobra.MarkFlagRequired(cmd.Flags(), "batch-key")

	cmd.Flags().String(
		"batch-name", "",
		"Azure batch account name",
	)
	cobra.MarkFlagRequired(cmd.Flags(), "batch-name")

	cmd.Flags().String(
		"storage-key", "",
		"Azure blob storage account key",
	)
	cobra.MarkFlagRequired(cmd.Flags(), "storage-key")

	cmd.Flags().String(
		"storage-name", "",
		"Azure blob storage account name",
	)
	cobra.MarkFlagRequired(cmd.Flags(), "storage-name")
}
