package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"tower-cli-go/internal"
	"tower-cli-go/internal/commands/common_flags"
	"tower-cli-go/internal/formatters"

	"github.com/spf13/cobra"
)

func NewLeaveWorkspacesCmd() *cobra.Command {

	leaveCmd := &cobra.Command{
		Use:   "leave",
		Short: "Leave workspace",
		RunE:  leaveRunE,
	}

	common_flags.AddWorkspaceRefFlag(leaveCmd)

	return leaveCmd
}

func leaveRunE(cmd *cobra.Command, args []string) error {

	wrapper := internal.NewApiFor(cmd)

	wspID, _ := cmd.Flags().GetInt64("id")
	wspName, _ := cmd.Flags().GetString("name")

	dto, err := wrapper.FetchOrgAndWspDbDto(wspID, wspName)
	if err != nil {
		return err
	}

	_, err = wrapper.Api.LeaveWorkspaceParticipant(wrapper.Ctx, dto.OrgId, dto.WorkspaceId)
	if err != nil {
		return err
	}

	result := WorkspaceLeave{
		WspName: dto.WorkspaceName,
	}

	return formatters.PrintTo(cmd.OutOrStdout(), result)
}

type WorkspaceLeave struct {
	WspName string `json:"workspaceName"`
}

func (res WorkspaceLeave) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(res)
}

func (res WorkspaceLeave) WriteAsTable(w io.Writer) error {
	_, err := fmt.Fprintf(w, "You have been removed as participant from '%s' workspace\n", res.WspName)
	return err
}
