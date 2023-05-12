package organizations

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization",
		RunE:  deleteRunE,
	}

	common_flags.AddOrganizationRefFlag(deleteCmd)

	return deleteCmd
}

func deleteRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	var orgRef string
	orgId, _ := cmd.Flags().GetInt64("id")
	orgName, _ := cmd.Flags().GetString("name")

	if orgId >= 0 {
		orgRef = strconv.Itoa(int(orgId))
	} else {
		orgAndWsp, err := wrapper.OrgByName(orgName)
		if err != nil {
			return err
		}
		orgId = orgAndWsp.OrgId
		orgRef = orgAndWsp.OrgName
	}

	_, err := wrapper.Api.DeleteOrganization(wrapper.Ctx, orgId)
	if err != nil {
		return fmt.Errorf("organization '%s' could not be deleted: %w", orgRef, err)
	}

	result := OrganizationsDeleted{orgRef}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type OrganizationsDeleted struct {
	Ref string `json: "ref"`
}

func (res OrganizationsDeleted) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(res)
}

func (res OrganizationsDeleted) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "Organization '%s' deleted\n", res.Ref)
	return err
}
