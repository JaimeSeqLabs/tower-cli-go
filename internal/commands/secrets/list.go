package secrets

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
		Short: "List workspace secrets",
		RunE: func(cmd *cobra.Command, args []string) error {

			wrapper := internal.NewApiFor(cmd)
			wspRefFlag, _ := cmd.Flags().GetString("workspace")

			wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
			if err != nil {
				return err
			}

			if !wspID.IsSet() {
				wspRef = "user"
			}

			response, httpRes, err := wrapper.Api.ListPipelineSecrets(
				wrapper.Ctx,
				&openapi.DefaultApiListPipelineSecretsOpts{
					WorkspaceId: wspID,
				},
			)
			if err != nil {
				if httpRes.StatusCode == 404 {
					return fmt.Errorf("workspace %d not found", wspID.Value())
				}
			}

			result := SecretsList{
				WorkspaceRef: wspRef,
				Secrets:      response.PipelineSecrets,
			}

			return formatters.PrintTo(cmd.OutOrStdout(), result)
		},
	}

	common_flags.AddOptionalWorkspaceFlags(listCmd)

	return listCmd
}

type SecretsList struct {
	WorkspaceRef string                   `json:"workspaceRef"`
	Secrets      []openapi.PipelineSecret `json:"secrets"`
}

func (list SecretsList) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(list)
}

func (list SecretsList) WriteAsTable(w io.Writer) error {

	fmt.Fprintf(w, "\nSecrets at workspace %s :\n\n", list.WorkspaceRef)

	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	if len(list.Secrets) == 0 {
		t.AppendRow(table.Row{"No secrets found"})
	} else {

		t.SortBy([]table.SortBy{
			{Name: "ID", Mode: table.Asc},
		})
		t.AppendHeader(table.Row{
			"ID", "Name", "Created", "Updated", "Used",
		})
		for _, s := range list.Secrets {

			var id string
			if s.Id != nil {
				id = fmt.Sprintf("%d", *s.Id)
			} else {
				id = "null"
			}

			t.AppendRow(table.Row{
				id,
				s.Name,
				formatters.FormatDate(s.DateCreated),
				formatters.FormatDate(s.LastUpdated),
				formatters.FormatDate(s.LastUsed),
			})

		}

	}

	t.Render()
	return nil
}
