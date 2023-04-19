package credentials

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
		Short: "Delete workspace credentials",
		RunE:  deleteRunE,
	}

	common_flags.AddOptionalWorkspaceFlags(deleteCmd)
	common_flags.AddOptionalCredentialsRefFlags(deleteCmd)

	return deleteCmd
}

func deleteRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspRefFlag, _ := cmd.Flags().GetString("workspace")
	credsId, _ := cmd.Flags().GetString("id")
	credsName, _ := cmd.Flags().GetString("name")

	wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
	if err != nil {
		return err
	}

	if !wspID.IsSet() {
		wspRef = "user"
	}

	var id string
	if credsId != "" {
		id = credsId
	} else {
		creds, _, err := wrapper.FindCredentialsByName(wspID, credsName)
		if err != nil {
			return err
		}
		id = creds.Id
	}

	httpRes, err := wrapper.Api.DeleteCredentials(wrapper.Ctx, id, &openapi.DefaultApiDeleteCredentialsOpts{WorkspaceId: wspID})
	if err != nil {
		if httpRes.StatusCode == 403 { // forbidden
			return fmt.Errorf("credentials '%s' not found in '%s' workspace", id, wspRef)
		}
		return err
	}

	result := CredentialsDeleted{
		ID:           id,
		WorkspaceRef: wspRef,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type CredentialsDeleted struct {
	ID           string `json:"id"`
	WorkspaceRef string `json:"workspaceRef"`
}

func (cd CredentialsDeleted) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(cd)
}

func (cd CredentialsDeleted) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "Credentials '%s' deleted at %s workspace\n", cd.ID, cd.WorkspaceRef)
	return err
}
