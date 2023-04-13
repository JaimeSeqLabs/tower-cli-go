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
// ResourceData struct for ResourceData
type ResourceData struct {
	Warnings []string `json:"warnings,omitempty"`
	Mean float32 `json:"mean,omitempty"`
	Min float32 `json:"min,omitempty"`
	Q1 float32 `json:"q1,omitempty"`
	Q2 float32 `json:"q2,omitempty"`
	Q3 float32 `json:"q3,omitempty"`
	Max float32 `json:"max,omitempty"`
	MinLabel string `json:"minLabel,omitempty"`
	MaxLabel string `json:"maxLabel,omitempty"`
	Q1Label string `json:"q1Label,omitempty"`
	Q2Label string `json:"q2Label,omitempty"`
	Q3Label string `json:"q3Label,omitempty"`
}