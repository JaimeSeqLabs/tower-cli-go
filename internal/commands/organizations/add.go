package organizations

import (
	"encoding/json"
	"fmt"
	"io"
	"openapi"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new organization",
		RunE:  addCmdRunE,
	}

	common_flags.AddOrganizationDataFlags(addCmd)

	addCmd.Flags().StringP(
		"name", "n", "",
		"Organization name",
	)
	cobra.MarkFlagRequired(addCmd.Flags(), "name")

	addCmd.Flags().StringP(
		"full-name", "f", "",
		"Organization full name",
	)
	cobra.MarkFlagRequired(addCmd.Flags(), "full-name")

	return addCmd
}

func addCmdRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	orgName, _ := cmd.Flags().GetString("name")
	orgFullName, _ := cmd.Flags().GetString("full-name")
	orgDesc, _ := cmd.Flags().GetString("description")
	orgLoc, _ := cmd.Flags().GetString("location")
	orgURL, _ := cmd.Flags().GetString("website")

	org := openapi.Organization{
		Name:        orgName,
		FullName:    orgFullName,
		Description: orgDesc,
		Location:    orgLoc,
		Website:     orgURL,
	}

	response, _, err := wrapper.Api.CreateOrganization(
		wrapper.Ctx,
		openapi.CreateOrganizationRequest{Organization: org},
	)
	if err != nil {
		return err
	}

	result := OrganizationsAdded{response.Organization}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type OrganizationsAdded struct {
	Organization openapi.OrganizationDbDto `json:"organization"`
}

func (response OrganizationsAdded) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(response)
}

func (response OrganizationsAdded) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "Organization '%s' with id '%d' was added\n", response.Organization.Name, response.Organization.OrgId)
	return err
}
