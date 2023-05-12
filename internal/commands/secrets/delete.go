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

func NewDeleteCmd() *cobra.Command {

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a workspace secret",
		RunE:  deleteRunE,
	}

	common_flags.AddOptionalWorkspaceFlags(deleteCmd)
	common_flags.AddSecretRefFlags(deleteCmd)

	return deleteCmd
}

func deleteRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspRefFlag, _ := cmd.Flags().GetString("workspace")
	secretID, _ := cmd.Flags().GetInt64("id")
	secretName, _ := cmd.Flags().GetString("name")

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

	_, err = wrapper.Api.DeletePipelineSecret(
		wrapper.Ctx,
		*secretResponse.Id,
		&openapi.DefaultApiDeletePipelineSecretOpts{
			WorkspaceId: wspID,
		},
	)
	if err != nil {
		return err
	}

	result := SecretDelete{
		WorkspaceRef: wspRef,
		Secret:       secretResponse,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type SecretDelete struct {
	WorkspaceRef string                 `json:"workspaceRef"`
	Secret       openapi.PipelineSecret `json:"secret"`
}

func (s SecretDelete) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(s)
}

func (s SecretDelete) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "Secret '%s' deleted at %s workspace\n", s.Secret.Name, s.WorkspaceRef)
	return err

}
