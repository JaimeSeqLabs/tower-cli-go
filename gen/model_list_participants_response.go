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
// ListParticipantsResponse struct for ListParticipantsResponse
type ListParticipantsResponse struct {
	TotalSize int64 `json:"totalSize"`
	Participants []ParticipantDbDto `json:"participants"`
}