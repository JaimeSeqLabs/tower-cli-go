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
// AssociatePipelineLabelsRequest struct for AssociatePipelineLabelsRequest
type AssociatePipelineLabelsRequest struct {
	PipelineIds []int64 `json:"pipelineIds,omitempty"`
	LabelIds []int64 `json:"labelIds,omitempty"`
}
