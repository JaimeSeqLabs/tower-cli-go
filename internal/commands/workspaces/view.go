package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/formatters"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func NewViewWorkspacesCmd() *cobra.Command {

	viewCmd := &cobra.Command{
		Use:   "view",
		Short: "Describe an existing organization workspace",
		RunE:  viewRunE,
	}

	common_flags.AddWorkspaceRefFlag(viewCmd)

	return viewCmd
}

func viewRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspID, _ := cmd.Flags().GetInt64("id")
	wspName, _ := cmd.Flags().GetString("name")

	dto, err := wrapper.FetchOrgAndWspDbDto(wspID, wspName)
	if err != nil {
		return err
	}

	res, _, err := wrapper.Api.DescribeWorkspace(wrapper.Ctx, dto.OrgId, dto.WorkspaceId)
	if err != nil {
		return err
	}

	result := WorkspaceView{
		Workspace: res.Workspace,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type WorkspaceView struct {
	Workspace openapi.Workspace `json:"workspace"`
}

func (res WorkspaceView) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(res)
}

func (res WorkspaceView) WriteAsTable(w io.Writer) error {

	fmt.Fprintf(w, "\nDetails for workspace '%s' :\n\n", res.Workspace.FullName)

	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	t.AppendHeader(table.Row{
		"ID", "Name", "Full Name", "Description", "Visibility", "Date Created", "Last Updated",
	})

	id := "null"
	if res.Workspace.Id != nil {
		id = fmt.Sprintf("%d", *res.Workspace.Id)
	}

	t.AppendRow(table.Row{
		id,
		res.Workspace.Name,
		res.Workspace.FullName,
		res.Workspace.Description,
		res.Workspace.Visibility,
		formatters.FormatDate(res.Workspace.DateCreated),
		formatters.FormatDate(res.Workspace.LastUpdated),
	})

	t.Render()
	return nil
}
