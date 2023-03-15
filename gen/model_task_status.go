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
// TaskStatus the model 'TaskStatus'
type TaskStatus string

// List of TaskStatus
const (
	NEW TaskStatus = "NEW"
	SUBMITTED TaskStatus = "SUBMITTED"
	RUNNING TaskStatus = "RUNNING"
	CACHED TaskStatus = "CACHED"
	COMPLETED TaskStatus = "COMPLETED"
	FAILED TaskStatus = "FAILED"
	ABORTED TaskStatus = "ABORTED"
)