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
// ActionConfigType struct for ActionConfigType
type ActionConfigType struct {
	Discriminator string `json:"discriminator,omitempty"`
	Events []string `json:"events,omitempty"`
}