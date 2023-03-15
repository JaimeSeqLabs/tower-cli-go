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
// OrgRole the model 'OrgRole'
type OrgRole string

// List of OrgRole
const (
	OWNER OrgRole = "owner"
	MEMBER OrgRole = "member"
	COLLABORATOR OrgRole = "collaborator"
)