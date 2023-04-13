package organizations

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	. "tower-cli-go/internal"
	"tower-cli-go/internal/formatters"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)


var ViewCmd = &cobra.Command{
	Use: "view",
	Short: "Describe organization details",
	RunE: func(cmd *cobra.Command, args []string) error {

		orgId, _ := cmd.Flags().GetInt64("id")
		orgName, _ := cmd.Flags().GetString("name")
		
		wrapper := NewApiFor(cmd)

		response, err := wrapper.FetchOrganization(orgId, orgName)
		if err != nil {
			return err
		}

		result := OrganizationsView {
			Organization: response.Organization,
			ServerUrl: wrapper.ServerUrl(),
		}

		return formatters.PrintTo(cmd.OutOrStdout(), result)
	},
}

var (
	_orgID int64
	_orgName string
)

func init() {
	ViewCmd.Flags().Int64VarP(&_orgID, "id", "i", 0, "Organization unique id")
	ViewCmd.Flags().StringVarP(&_orgName, "name", "n", "", "Organization unique id")
	ViewCmd.MarkFlagsMutuallyExclusive("id", "name")
}

type OrganizationsView struct {
	Organization openapi.OrganizationDbDto `json:"organization"`
	ServerUrl string `json:"-"`
}

func (response OrganizationsView) WriteAsTable(w io.Writer) error {
	
	fmt.Fprintf(w, "\nDetails for organization %s :\n\n", response.Organization.FullName)
	
	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	t.AppendRow(table.Row{
		"ID", formatters.FormatOrgID(response.Organization.OrgId, response.ServerUrl, response.Organization.Name),
	})
	t.AppendRow(table.Row{
		"Name", response.Organization.Name,
	})
	t.AppendRow(table.Row{
		"Full Name", response.Organization.FullName,
	})
	t.AppendRow(table.Row{
		"Description", response.Organization.Description,
	})
	t.AppendRow(table.Row{
		"Website", response.Organization.Website,
	})

	t.Render()
	return nil
}

func (response OrganizationsView) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(response)
}