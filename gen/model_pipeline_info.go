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
// PipelineInfo struct for PipelineInfo
type PipelineInfo struct {
	ProjectName string `json:"projectName,omitempty"`
	SimpleName string `json:"simpleName,omitempty"`
	RepositoryUrl string `json:"repositoryUrl,omitempty"`
	CloneUrl string `json:"cloneUrl,omitempty"`
	Provider string `json:"provider,omitempty"`
	ConfigFiles []string `json:"configFiles,omitempty"`
	WorkDirs []string `json:"workDirs,omitempty"`
	Revisions []string `json:"revisions,omitempty"`
	Profiles []string `json:"profiles,omitempty"`
	Manifest WfManifest `json:"manifest,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}