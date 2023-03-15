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
// TraceCreateRequest struct for TraceCreateRequest
type TraceCreateRequest struct {
	SessionId string `json:"sessionId,omitempty"`
	RunName string `json:"runName,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	Repository string `json:"repository,omitempty"`
	WorkflowId string `json:"workflowId,omitempty"`
}
