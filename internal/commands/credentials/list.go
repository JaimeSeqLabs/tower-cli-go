package credentials

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

func NewListCmd() *cobra.Command {

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all worskpace credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			wrapper := internal.NewApiFor(cmd)
			wspRefFlag, _ := cmd.Flags().GetString("workspace")

			wspID, _, _, wspRef, _ := wrapper.WorkspaceIdentifiers(wspRefFlag)

			if !wspID.IsSet() {
				wspRef = "user"
			}

			response, _, err := wrapper.Api.ListCredentials(wrapper.Ctx, &openapi.DefaultApiListCredentialsOpts{WorkspaceId: wspID})
			if err != nil {
				return err
			}

			wspURL, err := wrapper.BaseWspUrl(wspID.Value())
			if err != nil {
				return err
			}

			// response.Credentials is "omitIfEmpty"
			creds := response.Credentials
			if creds == nil {
				creds = []openapi.Credentials{}
			}

			result := CredentialsList{
				WorskpaceRef: wspRef,
				Credentials:  creds,
				BaseWspUrl:   wspURL,
			}

			return formatters.PrintTo(cmd.OutOrStdout(), result)
		},
	}

	common_flags.AddOptionalWorkspaceFlags(listCmd)

	return listCmd
}

type CredentialsList struct {
	WorskpaceRef string                `json:"workspaceRef"`
	Credentials  []openapi.Credentials `json:"credentials"`
	BaseWspUrl   string                `json:"-"`
}

func (response CredentialsList) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(response)
}

func (response CredentialsList) WriteAsTable(w io.Writer) error {

	fmt.Fprintf(w, "\nCredentials at workspace %s :\n\n", response.WorskpaceRef)

	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	if len(response.Credentials) == 0 {
		t.AppendRow(table.Row{"No credentials found"})

	} else {
		t.SortBy([]table.SortBy{
			{Name: "ID", Mode: table.Asc},
		})
		t.AppendHeader(table.Row{
			"ID", "Provider", "Name", "Last activity",
		})
		for _, cred := range response.Credentials {
			t.AppendRow(table.Row{
				formatters.FormatCredentialsID(cred.Id, response.BaseWspUrl),
				cred.Provider,
				cred.Name,
				formatters.FormatTime(cred.LastUsed),
			})
		}
	}

	t.Render()
	return nil
}
