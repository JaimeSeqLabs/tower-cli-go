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
// MemberDbDto struct for MemberDbDto
type MemberDbDto struct {
	MemberId int64 `json:"memberId,omitempty"`
	Email string `json:"email,omitempty"`
	UserName string `json:"userName,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Role OrgRole `json:"role,omitempty"`
}
