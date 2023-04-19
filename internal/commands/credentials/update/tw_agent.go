package update_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewUpdateTwAgentCmd() *cobra.Command {

	twaCmd := &cobra.Command{
		Use:   "agent",
		Short: "Update Tower Agent workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewTwAgentCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(twaCmd)

	twaCmd.Flags().String(
		"connection-id", "",
		"Connection identifier",
	)
	cobra.MarkFlagRequired(twaCmd.Flags(), "connection-id")

	twaCmd.Flags().String(
		"work-dir", "",
		"Default work directory",
	)
	viper.BindEnv("work-dir", "TW_AGENT_WORK")

	return twaCmd
}
