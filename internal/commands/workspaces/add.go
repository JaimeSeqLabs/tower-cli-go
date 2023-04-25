package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"strings"
	"tower-cli-go/internal"
	"tower-cli-go/internal/formatters"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

func NewAddWorkspacesCmd() *cobra.Command {

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new organization workspace",
		RunE:  addRunE,
	}

	addCmd.Flags().String(
		"org", "",
		"The workspace organization name",
	)

	addCmd.Flags().StringP(
		"name", "n", "",
		"The workspace short name. Only alphanumeric, dash and underscore characters are allowed",
	)

	addCmd.Flags().StringP(
		"full-name", "f", "",
		"The workspace full name",
	)

	addCmd.MarkFlagRequired("organization")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("full-name")

	addCmd.Flags().StringP(
		"description", "d", "",
		"The workspace description",
	)

	// TODO: -v shorthand collides with global -v verbose flag
	addCmd.Flags().String(
		"visibility", "PRIVATE",
		"The workspace visibility. Valid options PRIVATE, SHARED",
	)

	return addCmd
}

func addRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	orgName, _ := cmd.Flags().GetString("organization")
	wspName, _ := cmd.Flags().GetString("name")
	wspFullName, _ := cmd.Flags().GetString("full-name")
	wspDesc, _ := cmd.Flags().GetString("description")
	wspVisFlag, _ := cmd.Flags().GetString("visibility")

	wspVis, err := parseVisibility(wspVisFlag)
	if err != nil {
		return err
	}

	dto, err := wrapper.OrgByName(orgName)
	if err != nil {
		return err
	}

	workspace := openapi.Workspace{
		Name:        wspName,
		FullName:    wspFullName,
		Description: wspDesc,
		Visibility:  wspVis,
	}

	_, err = wrapper.Api.WorkspaceValidate(wrapper.Ctx, dto.OrgId, &openapi.DefaultApiWorkspaceValidateOpts{Name: optional.NewString(wspName)})
	if err != nil {
		return err
	}

	response, _, err := wrapper.Api.CreateWorkspace(wrapper.Ctx, dto.OrgId, openapi.CreateWorkspaceRequest{Workspace: workspace})
	if err != nil {
		return err
	}

	result := WorkspaceAdded{
		WspName:    response.Workspace.Name,
		OrgName:    orgName,
		Visibility: response.Workspace.Visibility,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

func parseVisibility(vis string) (openapi.Visibility, error) {
	visUpper := strings.ToUpper(vis)
	switch visUpper {
	case "PRIVATE":
		return openapi.VISIBILITY_PRIVATE, nil
	case "SHARED":
		return openapi.VISIBILITY_SHARED, nil
	default:
		return "", fmt.Errorf("invalid visibility value, expected one of [PRIVATE, SHARED] but got '%s' instead", vis)
	}
}

type WorkspaceAdded struct {
	WspName    string             `json:"workspaceName"`
	OrgName    string             `json:"organizationName"`
	Visibility openapi.Visibility `json:"visibility"`
}

func (res WorkspaceAdded) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(res)
}

func (res WorkspaceAdded) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "A '%s' workspace '%s' added for '%s' organization\n", res.Visibility, res.WspName, res.OrgName)
	return err
}
