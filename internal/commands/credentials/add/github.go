package add_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddGitHubCmd() *cobra.Command {

	ghCmd := &cobra.Command{
		Use:   "github",
		Short: "Add new GitHub workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewGitHubCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(ghCmd)

	// TODO: shorthand -u collides with global -u url
	ghCmd.Flags().String(
		"username", "",
		"GitHub username",
	)
	cobra.MarkFlagRequired(ghCmd.Flags(), "username")

	ghCmd.Flags().StringP(
		"password", "p", "",
		"GitHub account password or access token (recommended)",
	)
	cobra.MarkFlagRequired(ghCmd.Flags(), "password")

	// TODO: -t shorthand collides with global -t tower api token
	//glabCmd.Flags().StringP("token", "t", "", "GitHub account password or access token (recommended)")
	ghCmd.Flags().String(
		"github-token", "",
		"GitHub account access token",
	)
	cobra.MarkFlagRequired(ghCmd.Flags(), "github-token")

	ghCmd.Flags().String(
		"base-url", "",
		"Repository base URL",
	)

	return ghCmd
}
