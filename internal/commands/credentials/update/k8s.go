package update_cmd

import (
	"tower-cli-go/internal/credentials/provider"

	"github.com/spf13/cobra"
)

func NewUpdateK8sCmd() *cobra.Command {

	k8sCmd := &cobra.Command{
		Use:   "k8s",
		Short: "Update Kubernetes workspace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			kProv := provider.NewK8sCredentialsProvider(cmd)
			runE := createUpdateCmdRunE(kProv)

			return runE(cmd, args)
		},
	}

	addCommonUpdateCmdFlags(k8sCmd)

	k8sCmd.Flags().StringP(
		"certificate", "c", "",
		"Client certificate file",
	)

	k8sCmd.Flags().StringP(
		"private-key", "k", "",
		"Client key file",
	)

	// TODO: -t shorthand collides with global -t tower api token
	//k8sCmd.Flags().StringP("token", "t", "", "Kubernetes token")
	k8sCmd.Flags().String(
		"k8s-token", "",
		"Kubernetes token",
	)

	return k8sCmd
}
