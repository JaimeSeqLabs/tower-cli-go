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
// DescribeUserResponse struct for DescribeUserResponse
type DescribeUserResponse struct {
	User User `json:"user,omitempty"`
	NeedConsent bool `json:"needConsent,omitempty"`
	DefaultWorkspaceId int64 `json:"defaultWorkspaceId,omitempty"`
}
