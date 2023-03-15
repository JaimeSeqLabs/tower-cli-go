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
// AddParticipantRequest struct for AddParticipantRequest
type AddParticipantRequest struct {
	MemberId int64 `json:"memberId,omitempty"`
	TeamId int64 `json:"teamId,omitempty"`
	UserNameOrEmail string `json:"userNameOrEmail,omitempty"`
}
