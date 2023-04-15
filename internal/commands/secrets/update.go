package secrets

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update a workspace secret",
		RunE: updateRunE,
	}

	common_flags.AddOptionalWorkspaceFlags(updateCmd)
	common_flags.AddSecretRefFlags(updateCmd)
	// shorthand -v collides with global -v verbose flag
	updateCmd.Flags().String("value", "", "Secret value")

	return updateCmd
}

func updateRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspRefFlag, _ := cmd.Flags().GetString("workspace")
	secretID, _ := cmd.Flags().GetInt64("id")
	secretName, _ := cmd.Flags().GetString("name")
	secretValue, _ := cmd.Flags().GetString("value")

	wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
	if err != nil {
		return err
	}

	if !wspID.IsSet() {
		wspRef = "user"
	}

	secretResponse, err := wrapper.FetchSecret(secretID, secretName, wspID)
	if err != nil {
		return err
	}

	_, err = wrapper.Api.UpdatePipelineSecret(
		wrapper.Ctx, 
		*secretResponse.Id, 
		openapi.UpdatePipelineSecretRequest{
			Value: secretValue,
		},
		&openapi.DefaultApiUpdatePipelineSecretOpts{
			WorkspaceId: wspID,
		},
	)
	if err != nil {
		return err
	}

	result := SecretUpdate{
		WorkspaceRef: wspRef,
		SecretName: secretResponse.Name,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type SecretUpdate struct {
	WorkspaceRef string `json:"workspaceRef"`
	SecretName string `json:"secretName"`
}

func (s SecretUpdate) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(s)
}

func (s SecretUpdate) WriteAsTable(w io.Writer) error {
	fmt.Fprintf(w, "Secret '%s' updated at %s workspace\n", s.SecretName, s.WorkspaceRef)
	return nil

}
