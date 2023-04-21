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

func NewUpdateCmd() *cobra.Command {

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update organization details",
		RunE:  updateCmdRunE,
	}

	common_flags.AddOrganizationRefFlag(updateCmd)
	common_flags.AddOrganizationDataFlags(updateCmd)
	updateCmd.Flags().StringP(
		"full-name", "f", "",
		"Organization full name",
	)

	return updateCmd
}

func updateCmdRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	orgID, _ := cmd.Flags().GetInt64("id")
	orgName, _ := cmd.Flags().GetString("name")
	orgFullName, _ := cmd.Flags().GetString("full-name")
	orgDesc, _ := cmd.Flags().GetString("description")
	orgLoc, _ := cmd.Flags().GetString("location")
	orgURL, _ := cmd.Flags().GetString("website")

	response, err := wrapper.FetchOrganization(orgID, orgName)
	if err != nil {
		return err
	}

	if orgFullName == "" {
		orgFullName = response.Organization.FullName
	}
	if orgDesc == "" {
		orgDesc = response.Organization.Description
	}
	if orgLoc == "" {
		orgLoc = response.Organization.Location
	}
	if orgURL == "" {
		orgURL = response.Organization.Website
	}

	req := openapi.UpdateOrganizationRequest{
		FullName:    orgFullName,
		Description: orgDesc,
		Location:    orgLoc,
		Website:     orgURL,
	}

	_, err = wrapper.Api.UpdateOrganization(wrapper.Ctx, response.Organization.OrgId, req)
	if err != nil {
		return err
	}

	result := OrganizationsUpdated{
		OrgID: response.Organization.OrgId,
		Name:  response.Organization.Name,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type OrganizationsUpdated struct {
	OrgID int64  `json:"orgId"`
	Name  string `json:"name"`
}

func (res OrganizationsUpdated) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(res)
}

func (response OrganizationsUpdated) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "Organization '%s' was updated\n", response.Name)
	return err
}
