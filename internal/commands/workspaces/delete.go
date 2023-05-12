package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewDeleteWorkspacesCmd() *cobra.Command {

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization workspace",
		RunE:  deleteRunE,
	}

	common_flags.AddWorkspaceRefFlag(deleteCmd)

	return deleteCmd
}

func deleteRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)
	wspID, _ := cmd.Flags().GetInt64("id")
	wspName, _ := cmd.Flags().GetString("name")

	dto, err := wrapper.FetchOrgAndWspDbDto(wspID, wspName) // wspName can be in "org/wsp" form
	if err != nil {
		return err
	}

	_, err = wrapper.Api.DeleteWorkspace(wrapper.Ctx, dto.OrgId, dto.WorkspaceId)
	if err != nil {
		return err
	}

	result := WorkspaceDeleted{WspRef: dto.WorkspaceName, OrgRef: dto.OrgName}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type WorkspaceDeleted struct {
	WspRef string `json:"workspaceRef"`
	OrgRef string `json:"organizationRef"`
}

func (result WorkspaceDeleted) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(result)
}

func (result WorkspaceDeleted) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "Workspace '%s' deleted for '%s' organization\n", result.WspRef, result.OrgRef)
	return err
}
