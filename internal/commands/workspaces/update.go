package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewUpdateWorkspacesCmd() *cobra.Command {

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing organization workspace",
		RunE:  updateRunE,
	}

	updateCmd.Flags().Int64P(
		"id", "i", -1,
		"Workspace ID to update",
	)
	updateCmd.MarkFlagRequired("id")

	updateCmd.Flags().StringP(
		"full-name", "f", "",
		"The workspace full name",
	)

	updateCmd.Flags().StringP(
		"description", "d", "",
		"The workspace description",
	)

	return updateCmd
}

func updateRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspID, _ := cmd.Flags().GetInt64("id")
	wspFullName, _ := cmd.Flags().GetString("full-name")
	wspDesc, _ := cmd.Flags().GetString("description")

	if wspFullName == "" && wspDesc == "" {
		return fmt.Errorf("required at least one option to update")
	}

	dto, err := wrapper.FindOrgAndWspByWspID(wspID)
	if err != nil {
		return err
	}

	wspDescription, _, err := wrapper.Api.DescribeWorkspace(wrapper.Ctx, dto.OrgId, dto.WorkspaceId)
	if err != nil {
		return err
	}

	req := openapi.UpdateWorkspaceRequest{
		FullName:    wspDescription.Workspace.FullName,
		Description: wspDescription.Workspace.Description,
	}

	if wspFullName != "" {
		req.FullName = wspFullName
	}

	if wspDesc != "" {
		req.FullName = wspDesc
	}

	req.Visibility = openapi.VISIBILITY_PRIVATE

	res, _, err := wrapper.Api.UpdateWorkspace(wrapper.Ctx, dto.OrgId, dto.WorkspaceId, req)
	if err != nil {
		return err
	}

	result := WorkspaceUpdated{
		WspName:    res.Workspace.Name,
		OrgName:    dto.OrgName,
		Visibility: res.Workspace.Visibility,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type WorkspaceUpdated struct {
	WspName    string             `json:"workspaceName"`
	OrgName    string             `json:"organizationName"`
	Visibility openapi.Visibility `json:"visibility"`
}

func (res WorkspaceUpdated) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(res)
}

func (res WorkspaceUpdated) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "A '%s' workspace '%s' updated for '%s' organization\n", res.Visibility, res.WspName, res.OrgName)
	return err
}
