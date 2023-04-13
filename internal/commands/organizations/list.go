package organizations

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal/formatters"
	. "tower-cli-go/internal"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)


var ListCmd = &cobra.Command{
	Use: "list",
	Short: "List available organizations",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		wrapper := NewApiFor(cmd)

		userId, userName, err := wrapper.GetUserInfo()
		if err != nil {
			return err
		}

		// retrieve wsp and orgs for user
		response, _, err := wrapper.Api.ListWorkspacesUser(cmd.Context(), userId)
		if err != nil {
			return err
		}

		// save only orgs from response
		orgs := []openapi.OrgAndWorkspaceDbDto{}
		
		for _, dto := range response.OrgsAndWorkspaces {
			if dto.WorkspaceId <= 0 {
				orgs = append(orgs, dto)
			}
		}

		result := OrganizationsList {
			UserName: userName,
			Organizations: orgs,
			ServerUrl: wrapper.ServerUrl(),
		}

		// print results
		return formatters.PrintTo(cmd.OutOrStdout(), result)
	},
}

type OrganizationsList struct {
	UserName string `json:"userName"`
	Organizations []openapi.OrgAndWorkspaceDbDto `json:"organizations"`
	ServerUrl string `json:"-"`
}

func (response OrganizationsList) WriteAsTable(w io.Writer) error {
	
	fmt.Fprintf(w, "\nOrganizations for user %s :\n\n", response.UserName)
	
	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	if len(response.Organizations) == 0 {
		t.AppendRow(table.Row{ "No organizations found" })

	} else {
		t.AppendHeader(table.Row{ 
			"ID", "Name",
		})
		for _, org := range response.Organizations {
			t.AppendRow(table.Row{
				formatters.FormatOrgID(org.OrgId, response.ServerUrl, org.OrgName),
				org.OrgName,
			})
		}
	}

	t.Render()
	return nil
}

func (response OrganizationsList) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(response)
}