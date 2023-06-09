/*
 * Nextflow Tower API
 *
 * Nextflow Tower service API
 *
 * API version: 1.0.0
 * Contact: info@seqera.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PipelineDbDto struct for PipelineDbDto
type PipelineDbDto struct {
	PipelineId int64 `json:"pipelineId,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon string `json:"icon,omitempty"`
	Repository string `json:"repository,omitempty"`
	UserId int64 `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	UserFirstName string `json:"userFirstName,omitempty"`
	UserLastName string `json:"userLastName,omitempty"`
	OrgId int64 `json:"orgId,omitempty"`
	OrgName string `json:"orgName,omitempty"`
	WorkspaceId int64 `json:"workspaceId,omitempty"`
	WorkspaceName string `json:"workspaceName,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Optimized bool `json:"optimized,omitempty"`
	Labels []LabelDbDto `json:"labels,omitempty"`
}
