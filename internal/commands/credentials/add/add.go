package add_cmd

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

func NewAddCmd() *cobra.Command {

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add new workspace credentials",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	addCmd.AddCommand(
		NewAddAwsCmd(),
		NewAddGoogleCmd(),
		NewAddGitHubCmd(),
		NewAddGitLabCmd(),
		NewAddBitBucketCmd(),
		NewAddSSHCmd(),
		NewAddK8sCmd(),
		NewAddAzureCmd(),
		NewAddTwAgentCmd(),
		NewAddContainerRegistryCmd(),
	)

	return addCmd
}

// addCommonAddCmdFlags sets the flags used by all 'add' subcommands
func addCommonAddCmdFlags(cmd *cobra.Command) {
	common_flags.AddOptionalWorkspaceFlags(cmd)
	cmd.Flags().StringP("name", "n", "", "Credentials name")
	cobra.MarkFlagRequired(cmd.Flags(), "name")
}

// createAddCmdRunE creates a RunE function from a CredentialsSpecProvider interface
// all the credentials add subcommands share the same run logic
func createAddCmdRunE(keyProvider provider.CredentialsProvider) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		wrapper := internal.NewApiFor(cmd)

		wspRefFlag, _ := cmd.Flags().GetString("workspace")
		name, _ := cmd.Flags().GetString("name")

		wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
		if err != nil {
			return err
		}

		keys, err := keyProvider.SecurityKeys()
		if err != nil {
			return err
		}

		spec := openapi.CredentialsSpec{
			Keys:     keys,
			Name:     name,
			BaseUrl:  keyProvider.BaseUrl(),
			Provider: string(keyProvider.Type()),
		}

		response, _, err := wrapper.Api.CreateCredentials(
			wrapper.Ctx,
			openapi.CreateCredentialsRequest{
				Credentials: spec,
			},
			&openapi.DefaultApiCreateCredentialsOpts{
				WorkspaceId: wspID,
			},
		)
		if err != nil {
			return err
		}

		result := CrendentialsAdded{
			ID:           response.CredentialsId,
			Provider:     spec.Provider,
			Name:         spec.Name,
			WorkspaceRef: wspRef,
		}

		return formatters.PrintTo(cmd.OutOrStdout(), result)
	}
}

type CrendentialsAdded struct {
	ID           string `json:"id"`
	Provider     string `json:"provider"`
	Name         string `json:"name"`
	WorkspaceRef string `json:"workspaceRef"`
}

func (ca CrendentialsAdded) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(ca)
}

func (ca CrendentialsAdded) WriteAsTable(w io.Writer) error {
	fmt.Fprintf(w, "New %s credentials '%s' (%s) added at %s workspace\n", ca.Provider, ca.Name, ca.ID, ca.WorkspaceRef)
	return nil
}
