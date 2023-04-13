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
// ActionStatus the model 'ActionStatus'
type ActionStatus string

// List of Action.Status
const (
	ACTIONSTATUS_CREATING ActionStatus = "CREATING"
	ACTIONSTATUS_ACTIVE ActionStatus = "ACTIVE"
	ACTIONSTATUS_ERROR ActionStatus = "ERROR"
	ACTIONSTATUS_PAUSED ActionStatus = "PAUSED"
)