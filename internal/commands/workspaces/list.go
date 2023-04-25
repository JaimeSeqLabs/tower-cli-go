package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal"
	"tower-cli-go/internal/formatters"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func NewListWorkspacesCmd() *cobra.Command {

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List user workspaces",
		RunE:  listRunE,
	}

	listCmd.Flags().String(
		"org", "",
		"The workspace organization name",
	)

	return listCmd
}

func listRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)
	orgName, _ := cmd.Flags().GetString("org")

	userID, userName, err := wrapper.GetUserInfo()
	if err != nil {
		return err
	}

	response, _, err := wrapper.Api.ListWorkspacesUser(wrapper.Ctx, userID)
	if err != nil {
		return err
	}

	result := WorkspaceList{
		UserName:   userName,
		ServerURL:  wrapper.ServerUrl(),
		Workspaces: []openapi.OrgAndWorkspaceDbDto{},
	}

	// empty list response
	if len(response.OrgsAndWorkspaces) == 0 {
		return formatters.PrintTo(cmd.OutOrStdout(), result)
	}

	for _, dto := range response.OrgsAndWorkspaces {
		if dto.WorkspaceId > 0 && (orgName == "" || orgName == dto.OrgName) {
			result.Workspaces = append(result.Workspaces, dto)
		}
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type WorkspaceList struct {
	UserName   string                         `json:"userName"`
	Workspaces []openapi.OrgAndWorkspaceDbDto `json:"workspaces"`
	ServerURL  string                         `json:"-"`
}

func (list WorkspaceList) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(list)
}

func (list WorkspaceList) WriteAsTable(w io.Writer) error {

	fmt.Fprintf(w, "\nWorkspaces for user '%s' :\n\n", list.UserName)

	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	if len(list.Workspaces) == 0 {
		t.AppendRow(table.Row{"No workspaces found"})
	} else {

		t.SortBy([]table.SortBy{
			{Name: "Workspace ID", Mode: table.Asc},
		})
		t.AppendHeader(table.Row{
			"Workspace ID", "Workspace Name", "Organization Name", "Organization ID",
		})
		for _, dto := range list.Workspaces {
			t.AppendRow(table.Row{
				formatters.FormatWspID(dto.WorkspaceId, list.ServerURL, dto.OrgName, dto.WorkspaceName),
				dto.WorkspaceName,
				dto.OrgName,
				formatters.FormatOrgID(dto.OrgId, list.ServerURL, dto.OrgName),
			})
		}

	}

	t.Render()
	return nil
}
