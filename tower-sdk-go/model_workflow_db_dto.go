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
// WorkflowDbDto struct for WorkflowDbDto
type WorkflowDbDto struct {
	CommandLine string `json:"commandLine,omitempty"`
	Params map[string]map[string]interface{} `json:"params,omitempty"`
	ProjectDir string `json:"projectDir,omitempty"`
	Profile string `json:"profile,omitempty"`
	Start time.Time `json:"start,omitempty"`
	OwnerId int64 `json:"ownerId,omitempty"`
	Repository string `json:"repository,omitempty"`
	Id string `json:"id,omitempty"`
	Submit time.Time `json:"submit,omitempty"`
	Complete time.Time `json:"complete,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	RunName string `json:"runName,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	WorkDir string `json:"workDir,omitempty"`
	CommitId string `json:"commitId,omitempty"`
	UserName string `json:"userName,omitempty"`
	ScriptId string `json:"scriptId,omitempty"`
	Revision string `json:"revision,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	ScriptName string `json:"scriptName,omitempty"`
	LaunchId string `json:"launchId,omitempty"`
	Status WorkflowStatus `json:"status,omitempty"`
	ConfigFiles []string `json:"configFiles,omitempty"`
	ConfigText string `json:"configText,omitempty"`
	Resume bool `json:"resume,omitempty"`
	Manifest WfManifest `json:"manifest,omitempty"`
	Nextflow WfNextflow `json:"nextflow,omitempty"`
	Stats WfStats `json:"stats,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorReport string `json:"errorReport,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	HomeDir string `json:"homeDir,omitempty"`
	Container string `json:"container,omitempty"`
	ContainerEngine string `json:"containerEngine,omitempty"`
	ScriptFile string `json:"scriptFile,omitempty"`
	LaunchDir string `json:"launchDir,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	ExitStatus int32 `json:"exitStatus,omitempty"`
	Success bool `json:"success,omitempty"`
}
