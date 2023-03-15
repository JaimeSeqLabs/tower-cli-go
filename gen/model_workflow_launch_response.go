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
import (
	"time"
)
// WorkflowLaunchResponse struct for WorkflowLaunchResponse
type WorkflowLaunchResponse struct {
	Id string `json:"id,omitempty"`
	ComputeEnv ComputeEnv `json:"computeEnv,omitempty"`
	Pipeline string `json:"pipeline,omitempty"`
	WorkDir string `json:"workDir,omitempty"`
	Revision string `json:"revision,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	ConfigProfiles []string `json:"configProfiles,omitempty"`
	UserSecrets []string `json:"userSecrets,omitempty"`
	WorkspaceSecrets []string `json:"workspaceSecrets,omitempty"`
	ConfigText string `json:"configText,omitempty"`
	TowerConfig string `json:"towerConfig,omitempty"`
	ParamsText string `json:"paramsText,omitempty"`
	PreRunScript string `json:"preRunScript,omitempty"`
	PostRunScript string `json:"postRunScript,omitempty"`
	MainScript string `json:"mainScript,omitempty"`
	EntryName string `json:"entryName,omitempty"`
	SchemaName string `json:"schemaName,omitempty"`
	Resume bool `json:"resume,omitempty"`
	PullLatest bool `json:"pullLatest,omitempty"`
	StubRun bool `json:"stubRun,omitempty"`
	ResumeDir string `json:"resumeDir,omitempty"`
	ResumeCommitId string `json:"resumeCommitId,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
}
