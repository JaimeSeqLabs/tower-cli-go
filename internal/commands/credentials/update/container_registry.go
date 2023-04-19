package update_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateContainerRegistryCmd() *cobra.Command {

	cRegCmd := &cobra.Command{
		Use:   "container-reg",
		Short: "Update Container Registry workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewContainerRegistryCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(cRegCmd)

	// TODO: shorthand -u collides with global -u url
	cRegCmd.Flags().String(
		"username", "",
		"The user name to grant you access to the container registry",
	)
	cobra.MarkFlagRequired(cRegCmd.Flags(), "username")

	cRegCmd.Flags().StringP(
		"password", "p", "",
		"The password to grant you access to the container registry",
	)
	cobra.MarkFlagRequired(cRegCmd.Flags(), "password")

	cRegCmd.Flags().StringP(
		"registry", "r", "docker.io",
		"The container registry server name",
	)

	return cRegCmd
}
