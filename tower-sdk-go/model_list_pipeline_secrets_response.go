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
// ListPipelineSecretsResponse struct for ListPipelineSecretsResponse
type ListPipelineSecretsResponse struct {
	PipelineSecrets []PipelineSecret `json:"pipelineSecrets,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
}
