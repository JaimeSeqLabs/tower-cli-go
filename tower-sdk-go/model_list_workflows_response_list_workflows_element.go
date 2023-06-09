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
// ListWorkflowsResponseListWorkflowsElement struct for ListWorkflowsResponseListWorkflowsElement
type ListWorkflowsResponseListWorkflowsElement struct {
	Labels []LabelDbDto `json:"labels,omitempty"`
	Optimized bool `json:"optimized,omitempty"`
	Starred bool `json:"starred,omitempty"`
	Workflow WorkflowDbDto `json:"workflow,omitempty"`
	Progress ProgressData `json:"progress,omitempty"`
}
