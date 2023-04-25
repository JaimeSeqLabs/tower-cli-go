package internal

import (
	"context"
	"fmt"
	"net/http"
	"openapi"
	"strconv"
	"strings"
	"tower-cli-go/internal/utils"

	"github.com/antihax/optional"
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

func (w ApiWrapper) FetchOrgAndWspDbDto(wspID int64, wspRef string) (openapi.OrgAndWorkspaceDbDto, error) {

	if wspID > 0 {
		return w.FindOrgAndWspByWspID(wspID)
	}
	if strings.Contains(wspRef, "/") {
		return w.FindOrgAndWspByWspRef(wspRef)
	}
	
	return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("invalid workspace namespace '%s'", wspRef)
}

func (w ApiWrapper) FetchSecret(secretID int64, secretName string, wspID optional.Int64) (openapi.PipelineSecret, error) {
	if secretID <= 0 {
		return w.SecretByName(wspID, secretName)
	} else {
		res, _, err := w.Api.DescribePipelineSecret(
			w.Ctx, 
			secretID, 
			&openapi.DefaultApiDescribePipelineSecretOpts{
				WorkspaceId: wspID,
			},
		)
		if err != nil {
			return openapi.PipelineSecret{}, err
		}
		return res.PipelineSecret, nil
	}
}

func (w ApiWrapper) FetchCredentials(credsID string, credsName string, wspID optional.Int64) (openapi.Credentials, *http.Response, error) {
	
	if credsID != "" {
	
		res, httpRes, err := w.Api.DescribeCredentials(
			w.Ctx,
			credsID,
			&openapi.DefaultApiDescribeCredentialsOpts{
				WorkspaceId: wspID,
			},
		)
		return res.Credentials, httpRes, err
	
	} else {

		res, httpRes, err := w.FindCredentialsByName(wspID, credsName)
		return res, httpRes, err

	}
}

func (w ApiWrapper) SecretByName(wspID optional.Int64, secretName string) (openapi.PipelineSecret, error) {
	
	res, _, err := w.Api.ListPipelineSecrets(
		w.Ctx,
		&openapi.DefaultApiListPipelineSecretsOpts{
			WorkspaceId: wspID,
		},
	)
	if err != nil {
		return openapi.PipelineSecret{}, err
	}

	for _, s := range res.PipelineSecrets {
		if secretName == s.Name {
			return s, nil
		}
	}

	ref, err := w.WorkspaceRef(wspID.Value())
	if err != nil {
		return openapi.PipelineSecret{}, err
	}

	return openapi.PipelineSecret{}, fmt.Errorf("secret %s not found in workspace %s", secretName, ref)
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

func (w ApiWrapper) FindOrgAndWspByName(orgName, wspName string) (openapi.OrgAndWorkspaceDbDto, error) {
	
	userId, userName, err := w.GetUserInfo()
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, err
	}

	wspAndOrgs, _, err := w.Api.ListWorkspacesUser(w.Ctx, userId)
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, err
	}

	if len(wspAndOrgs.OrgsAndWorkspaces) == 0 {
		return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("no organization or workspace found for user %s", userName)
	}

	// find dto with target org and wsp name
	for _, dto := range wspAndOrgs.OrgsAndWorkspaces {
		if dto.OrgName == orgName && dto.WorkspaceName == wspName {
			return dto, nil
		}
	}

	return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("org %s with wsp %s not found in user %s organizations", orgName, wspName, userName)
}

func (w ApiWrapper) FindOrgAndWspByWspID(wspID int64) (openapi.OrgAndWorkspaceDbDto, error) {

	userID, userName, err := w.GetUserInfo()
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, err
	}

	dtos, _, err := w.Api.ListWorkspacesUser(w.Ctx, userID)
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, err
	}

	if len(dtos.OrgsAndWorkspaces) == 0 {
		return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("no workspaces found for user")
	}

	// find target dto
	for _, orgAndWsp := range dtos.OrgsAndWorkspaces {
		if wspID == orgAndWsp.WorkspaceId {
			return orgAndWsp, nil
		}
	}

	return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("no workspace %d found in user %s workspaces", wspID, userName)
}

func (w ApiWrapper) FindOrgAndWspByWspRef(wspRef string) (openapi.OrgAndWorkspaceDbDto, error) {
	
	if wspRef == "" {
		return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("empty workspace reference")
	}

	// org/wsp ref
	if strings.Contains(wspRef, "/") {

		parts := strings.Split(strings.TrimSpace(wspRef), "/")
		if len(parts) != 2 {
			return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("invalid workspace reference %s", wspRef)
		}

		orgWspDto, err := w.FindOrgAndWspByName(parts[0], parts[1])
		if err != nil {
			return openapi.OrgAndWorkspaceDbDto{}, err
		}

		return orgWspDto, nil
	}

	// the ref is an id, parse as i64
	id, err := strconv.ParseInt(wspRef, 10 , 64)
	if err != nil {
		return openapi.OrgAndWorkspaceDbDto{}, fmt.Errorf("invalid workspace ID")
	}

	return w.FindOrgAndWspByWspID(id)
}

func (w ApiWrapper) FindCredentialsByName(wspID optional.Int64, credsName string) (openapi.Credentials, *http.Response, error) {

	credsList, httpRes, err := w.Api.ListCredentials(w.Ctx, &openapi.DefaultApiListCredentialsOpts{
		WorkspaceId: wspID,
		PlatformId: optional.EmptyString(),
	})
	if err != nil {
		return openapi.Credentials{}, httpRes, err
	}

	if len(credsList.Credentials) == 0{
		return openapi.Credentials{}, httpRes, fmt.Errorf("credentials not found for workspace '%d'", wspID.Value())
	}

	for _, c := range credsList.Credentials {
		if c.Name == credsName {
			return c, httpRes, nil
		}
	}

	return openapi.Credentials{}, httpRes, fmt.Errorf("credentials not found for workspace '%d'", wspID.Value())
}

func (w ApiWrapper) WorkspaceRef(wspID int64) (string, error) {

	if wspID <= 0 {
		return "user", nil
	}

	dto, err := w.FindOrgAndWspByWspID(wspID)
	if err != nil {
		return "", err
	}

	return w.BuildWorkspaceRef(dto.OrgName, dto.WorkspaceName), nil
}

func (w ApiWrapper) WorkspaceID(maybeEmptyWspRef string) (optional.Int64, error) {
	
	if maybeEmptyWspRef == "" {
		return optional.EmptyInt64(), nil
	} else {
	
		dto, err := w.FindOrgAndWspByWspRef(maybeEmptyWspRef)
		if err != nil {
			return optional.EmptyInt64(), err
		}
		return optional.NewInt64(dto.WorkspaceId), nil
	
	}
}

func (w ApiWrapper) BuildWorkspaceRef(orgName, wspName string) string {
	return fmt.Sprintf("[%s / %s]", orgName, wspName)	
}

func (w ApiWrapper) WorkspaceIdentifiers(wspIdOrRef string) (wspID optional.Int64, wspName string, orgName string, ref string, retErr error) {

	wspID = optional.EmptyInt64()
	wspName = ""
	orgName = ""
	ref = ""
	retErr = nil

	if wspIdOrRef == "" {
		return
	}
	
	dto, err := w.FindOrgAndWspByWspRef(wspIdOrRef)
	if err != nil {
		retErr = err
		return
	}

	wspID = optional.NewInt64(dto.WorkspaceId)
	wspName = dto.WorkspaceName
	orgName = dto.OrgName
	ref = w.BuildWorkspaceRef(orgName, wspName)
	retErr = nil
	
	return
}

func (w ApiWrapper) BaseWspUrl(wspID int64) (string, error) {

	_, userName, err := w.GetUserInfo()
	if err != nil {
		return "", err
	}

	if wspID <= 0 {
		return fmt.Sprintf("%s/user/%s", w.ServerUrl(), userName), nil
	}

	dto, err := w.FindOrgAndWspByWspID(wspID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/orgs/%s/workspaces/%s", w.ServerUrl(), dto.OrgName, dto.WorkspaceName), nil
}

