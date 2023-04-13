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
// EventType struct for EventType
type EventType struct {
	Source string `json:"source,omitempty"`
	Display string `json:"display,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}