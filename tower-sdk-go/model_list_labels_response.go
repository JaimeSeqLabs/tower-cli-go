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
// ListLabelsResponse struct for ListLabelsResponse
type ListLabelsResponse struct {
	Labels []LabelDbDto `json:"labels,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
}