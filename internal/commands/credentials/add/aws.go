package add_cmd

import (
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewAddAwsCmd() *cobra.Command {

	awsCmd := &cobra.Command{
		Use:   "aws",
		Short: "Add new AWS workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewAwsCredentialsProvider(cmd)
			runE := createAddCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonAddCmdFlags(awsCmd)

	common_flags.AddAwsKeysFlags(awsCmd)

	awsCmd.Flags().StringP(
		"assume-role-arn", "r", "",
		"The IAM role to access the AWS resources. It should be a fully qualified AWS role ARN",
	)

	return awsCmd
}
