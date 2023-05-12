package update_cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/credentials/provider"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update workspace credentials",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	updateCmd.AddCommand(
		NewUpdateAwsCmd(),
		NewUpdateGoogleCmd(),
		NewUpdateGitHubCmd(),
		NewUpdateGitLabCmd(),
		NewUpdateBitBucketCmd(),
		NewUpdateSSHCmd(),
		NewUpdateK8sCmd(),
		NewUpdateAzureCmd(),
		NewUpdateContainerRegistryCmd(),
		NewUpdateTwAgentCmd(),
	)

	return updateCmd
}

// addCommonUpdateCmdFlags sets the flags used by all 'update' subcommands
func addCommonUpdateCmdFlags(cmd *cobra.Command) {
	common_flags.AddOptionalWorkspaceFlags(cmd)
	common_flags.AddOptionalCredentialsRefFlags(cmd)
}

// createAddCmdRunE creates a RunE function from a CredentialsSpecProvider interface
// all the credentials 'update' subcommands share the same run logic
func createUpdateCmdRunE(keyProvider provider.CredentialsProvider) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		wrapper := internal.NewApiFor(cmd)

		wspRefFlag, _ := cmd.Flags().GetString("workspace")
		credsName, _ := cmd.Flags().GetString("name")
		credsId, _ := cmd.Flags().GetString("id")

		wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
		if err != nil {
			return err
		}

		// check creds exist
		credentials, httpRes, err := wrapper.FetchCredentials(credsId, credsName, wspID)
		if err != nil {
			if httpRes.StatusCode == 403 { // forbidden
				if credsId != "" {
					return fmt.Errorf("credentials '%s' not found in '%s' workspace", credsId, wspRef)
				}
				return fmt.Errorf("credentials '%s' not found in '%s' workspace", credsName, wspRef)
			}
		}

		// build key objects
		keys, err := keyProvider.SecurityKeys()
		if err != nil {
			return err
		}
		spec := openapi.CredentialsSpec{
			Keys:     keys,
			Name:     credentials.Name,
			BaseUrl:  keyProvider.BaseUrl(),
			Provider: string(keyProvider.Type()),
			Id:       credentials.Id,
		}

		// do update
		_, err = wrapper.Api.UpdateCredentials(
			wrapper.Ctx,
			credentials.Id,
			openapi.UpdateCredentialsRequest{
				Credentials: spec,
			},
			&openapi.DefaultApiUpdateCredentialsOpts{
				WorkspaceId: wspID,
			},
		)
		if err != nil {
			return err
		}

		result := CredentialsUpdated{}

		return formatters.PrintTo(cmd.OutOrStdout(), result)
	}
}

type CredentialsUpdated struct {
	Provider     string `json:"provider"`
	Name         string `json:"name"`
	WorkspaceRef string `json:"workspaceRef"`
}

func (cu CredentialsUpdated) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(cu)
}

func (cu CredentialsUpdated) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s credentials '%s' updated at %s workspace\n", cu.Provider, cu.Name, cu.WorkspaceRef)
	return err
}
