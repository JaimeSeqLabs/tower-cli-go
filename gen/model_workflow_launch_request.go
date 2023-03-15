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
// WorkflowLaunchRequest struct for WorkflowLaunchRequest
type WorkflowLaunchRequest struct {
	Id string `json:"id,omitempty"`
	ComputeEnvId string `json:"computeEnvId,omitempty"`
	RunName string `json:"runName,omitempty"`
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
	OptimizationId string `json:"optimizationId,omitempty"`
	LabelIds []int64 `json:"labelIds,omitempty"`
	HeadJobCpus int32 `json:"headJobCpus,omitempty"`
	HeadJobMemoryMb int32 `json:"headJobMemoryMb,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
}
