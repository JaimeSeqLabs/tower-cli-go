package update_cmd

import (
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateAwsCmd() *cobra.Command {

	awsCmd := &cobra.Command{
		Use:   "aws",
		Short: "Update AWS workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewAwsCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(awsCmd)

	common_flags.AddAwsKeysFlags(awsCmd)

	awsCmd.Flags().StringP(
		"assume-role-arn", "r", "",
		"The IAM role to access the AWS resources. It should be a fully qualified AWS role ARN",
	)

	return awsCmd
}
