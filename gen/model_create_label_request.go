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
// CreateLabelRequest struct for CreateLabelRequest
type CreateLabelRequest struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Resource bool `json:"resource,omitempty"`
}
