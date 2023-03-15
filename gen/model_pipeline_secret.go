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
import (
	"time"
)
// PipelineSecret struct for PipelineSecret
type PipelineSecret struct {
	Id *int64 `json:"id,omitempty"`
	Name string `json:"name"`
	LastUsed time.Time `json:"lastUsed,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
}
