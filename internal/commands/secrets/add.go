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

func NewAddCmd() *cobra.Command {

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a workspace secret",
		RunE: func(cmd *cobra.Command, args []string) error {

			wrapper := internal.NewApiFor(cmd)

			wspRefFlag, _ := cmd.Flags().GetString("workspace")
			name, _ := cmd.Flags().GetString("name")
			value, _ := cmd.Flags().GetString("value")

			wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
			if err != nil {
				return err
			}

			if !wspID.IsSet() {
				wspRef = "user"
			}

			response, _, err := wrapper.Api.CreatePipelineSecret(
				wrapper.Ctx,
				openapi.CreatePipelineSecretRequest{
					Name: name,
					Value: value,
				},
				&openapi.DefaultApiCreatePipelineSecretOpts{
					WorkspaceId: wspID,
				},
			)
			if err != nil {
				return err
			}

			result := SecretAdded{
				WorkspaceRef: wspRef,
				ID: response.SecretId,
				Name: name,
			}

			return formatters.PrintTo(cmd.OutOrStdout(), result)
		},
	}

	// cannot have -v shorthand, collision with global -v 'verbose'
	addCmd.Flags().String("name", "", "Secret name")
	addCmd.Flags().String("value", "", "Secret value")
	cobra.MarkFlagRequired(addCmd.Flags(), "name")

	common_flags.AddOptionalWorkspaceFlags(addCmd)

	return addCmd
}

type SecretAdded struct {
	WorkspaceRef string `json:"workspaceRef"`
	ID int64 `json:"id"`
	Name string `json:"name"`
}

func (s SecretAdded) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(s)
}

func (s SecretAdded) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "New secret '%s' (%d) added at %s workspace\n", s.Name, s.ID, s.WorkspaceRef)
	return err
}
