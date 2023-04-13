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
// RunRequest struct for RunRequest
type RunRequest struct {
	WorkflowParams map[string]interface{} `json:"workflow_params,omitempty"`
	WorkflowType string `json:"workflow_type,omitempty"`
	WorkflowTypeVersion string `json:"workflow_type_version,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	WorkflowEngineParameters map[string]string `json:"workflow_engine_parameters,omitempty"`
	WorkflowUrl string `json:"workflow_url,omitempty"`
}