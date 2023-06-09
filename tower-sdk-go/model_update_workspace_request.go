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
// UpdateWorkspaceRequest struct for UpdateWorkspaceRequest
type UpdateWorkspaceRequest struct {
	Visibility Visibility `json:"visibility,omitempty"`
	Description string `json:"description"`
	FullName string `json:"fullName"`
}
