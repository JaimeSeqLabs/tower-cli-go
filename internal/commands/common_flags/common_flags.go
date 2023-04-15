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