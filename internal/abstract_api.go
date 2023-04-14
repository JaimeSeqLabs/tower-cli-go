package internal

import (
	"context"
	"fmt"
	"openapi"
	"strings"
	"tower-cli-go/internal/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ApiWrapper struct {
	Cmd *cobra.Command
	Api *openapi.DefaultApiService
	Ctx context.Context
}

func NewWrapperFrom(cmd *cobra.Command, api *openapi.DefaultApiService) ApiWrapper {
	return ApiWrapper{
		Cmd: cmd,
		Api: api,
		Ctx: cmd.Context(),
	}
}

func NewApiFor(cmd *cobra.Command) ApiWrapper {
	api := utils.GenerateClientFromCfg().DefaultApi
	return ApiWrapper{
		Cmd: cmd,
		Api: api,
		Ctx: cmd.Context(),
	}
}

func (w ApiWrapper) ServerUrl() string {
	url := w.ApiUrl()
	url = strings.Replace(url, "api.", "", 1)
	url = strings.Replace(url, "/api", "", 1)
	return url
}

func (w ApiWrapper) ApiUrl() string {
	if w.Cmd == nil {
		return viper.GetString("url")
	}
	url, _ := w.Cmd.Flags().GetString("url")
	return url
}

func (w ApiWrapper) GetUserInfo() (id int64, name string, err error) {
	
	name = ""
	id = -1
	err = nil

	userInfo, _, apiErr := w.Api.UserInfo(w.Ctx)
	if apiErr != nil {
		err = apiErr
		return
	}

	if userInfo.User.Id != nil {
		id = *userInfo.User.Id
	}

	name = userInfo.User.UserName
	
	return
}

func (w ApiWrapper) FetchOrganization(orgID int64, orgName string) (openapi.DescribeOrganizationResponse, error) {

	if orgID <= 0 { // no ID, find org dto by name then use its ID
		
		if orgName == "" {
			return openapi.DescribeOrganizationResponse{}, fmt.Errorf("organization ID or name needed")
		}

		dto, err := w.OrgByName(orgName)
		if err != nil {
			return openapi.DescribeOrganizationResponse{}, err
		}
		orgID = dto.OrgId
	}

	response, _, err := w.Api.DescribeOrganization(w.Ctx, orgID)
	if err != nil {
		return openapi.DescribeOrganizationResponse{}, err
	}
	
	return response, nil
}

func (w ApiWrapper) OrgByName(orgName string) (openapi.OrgAndWorkspaceDbDto, error) {

	userID, _, err := w.GetUserInfo()
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, err
	}

	orgAndWsps, _, err := w.Api.ListWorkspacesUser(w.Ctx, userID)
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, err
	}

	if len(orgAndWsps.OrgsAndWorkspaces) == 0 {
		return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("organization %s not found", orgName)
	}

	// get only orgs
	orgs := []openapi.OrgAndWorkspaceDbDto{}
	for _, e := range orgAndWsps.OrgsAndWorkspaces {
		if e.OrgName == orgName {
			orgs = append(orgs, e)
		}
	}

	if len(orgs) == 0 {
		return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("organization %s not found", orgName)
	}

	return orgs[1], nil
}
