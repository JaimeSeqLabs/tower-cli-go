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
// Workflow struct for Workflow
type Workflow struct {
	Status WorkflowStatus `json:"status,omitempty"`
	OwnerId int64 `json:"ownerId,omitempty"`
	Repository string `json:"repository,omitempty"`
	Id string `json:"id,omitempty"`
	Submit time.Time `json:"submit"`
	Start time.Time `json:"start,omitempty"`
	Complete time.Time `json:"complete,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	RunName string `json:"runName"`
	SessionId string `json:"sessionId"`
	Profile string `json:"profile,omitempty"`
	WorkDir string `json:"workDir"`
	CommitId string `json:"commitId,omitempty"`
	UserName string `json:"userName"`
	ScriptId string `json:"scriptId,omitempty"`
	Revision string `json:"revision,omitempty"`
	CommandLine string `json:"commandLine"`
	ProjectName string `json:"projectName"`
	ScriptName string `json:"scriptName,omitempty"`
	LaunchId string `json:"launchId,omitempty"`
	ConfigFiles []string `json:"configFiles,omitempty"`
	Params map[string]map[string]interface{} `json:"params,omitempty"`
	ConfigText string `json:"configText,omitempty"`
	Manifest WfManifest `json:"manifest,omitempty"`
	Nextflow WfNextflow `json:"nextflow,omitempty"`
	Stats WfStats `json:"stats,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorReport string `json:"errorReport,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	ProjectDir string `json:"projectDir,omitempty"`
	HomeDir string `json:"homeDir,omitempty"`
	Container string `json:"container,omitempty"`
	ContainerEngine string `json:"containerEngine,omitempty"`
	ScriptFile string `json:"scriptFile,omitempty"`
	LaunchDir string `json:"launchDir,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	ExitStatus int32 `json:"exitStatus,omitempty"`
	Resume bool `json:"resume,omitempty"`
	Success bool `json:"success,omitempty"`
	LogFile string `json:"logFile,omitempty"`
	OutFile string `json:"outFile,omitempty"`
	OperationId string `json:"operationId,omitempty"`
}
