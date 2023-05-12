package add_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddGitLabCmd() *cobra.Command {

	glabCmd := &cobra.Command{
		Use:   "gitlab",
		Short: "Add new GitLab workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewGitLabCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(glabCmd)

	// TODO: shorthand -u collides with global -u url
	glabCmd.Flags().String(
		"username", "",
		"GitLab username",
	)
	cobra.MarkFlagRequired(glabCmd.Flags(), "username")

	glabCmd.Flags().StringP(
		"password", "p", "",
		"GitLab account password or access token (recommended)",
	)
	cobra.MarkFlagRequired(glabCmd.Flags(), "password")

	// TODO: -t shorthand collides with global -t tower api token
	//glabCmd.Flags().StringP("token", "t", "", "GitLab account password or access token (recommended)")
	glabCmd.Flags().String(
		"gitlab-token", "",
		"GitLab account access token",
	)
	cobra.MarkFlagRequired(glabCmd.Flags(), "gitlab-token")

	glabCmd.Flags().String(
		"base-url", "",
		"Repository base URL",
	)

	return glabCmd
}
