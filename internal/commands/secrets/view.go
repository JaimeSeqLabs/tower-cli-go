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

func NewViewCmd() *cobra.Command {

	viewCmd := &cobra.Command{
		Use:   "view",
		Short: "View secret details",
		RunE: viewRunE,
	}

	common_flags.AddOptionalWorkspaceFlags(viewCmd)
	common_flags.AddSecretRefFlags(viewCmd)

	return viewCmd
}

func viewRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspRefFlag, _ := cmd.Flags().GetString("workspace")
	secretID, _ := cmd.Flags().GetInt64("id")
	secretName, _ := cmd.Flags().GetString("name")

	wspID, _, _, wspRef, err := wrapper.WorkspaceIdentifiers(wspRefFlag)
	if err != nil {
		return err
	}

	if !wspID.IsSet() {
		wspRef = "user"
	}

	response, err := wrapper.FetchSecret(secretID, secretName, wspID)
	if err != nil {
		return err
	}

	result := SecretView{
		WorkspaceRef: wspRef,
		Secret: response,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type SecretView struct {
	WorkspaceRef string `json:"workspaceRef"`
	Secret openapi.PipelineSecret `json:"secret"`
}

func (s SecretView) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(s)
}

func (s SecretView) WriteAsTable(w io.Writer) error {

	fmt.Fprintf(w, "\nSecret at workspace %s :\n\n", s.WorkspaceRef)
	
	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	var id string
	if s.Secret.Id != nil {
		id = fmt.Sprintf("%d", *s.Secret.Id)
	} else {
		id = "null"
	}

	t.AppendRows([]table.Row{
		{ "ID", id },
		{ "Name", s.Secret.Name },
		{ "Created", formatters.FormatDate(s.Secret.DateCreated) },
		{ "Updated", formatters.FormatDate(s.Secret.LastUpdated) },
		{ "Used", formatters.FormatDate(s.Secret.LastUsed) },
	})

	t.Render()
	return nil
}
