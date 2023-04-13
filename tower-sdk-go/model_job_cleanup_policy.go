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
// JobCleanupPolicy the model 'JobCleanupPolicy'
type JobCleanupPolicy string

// List of JobCleanupPolicy
const (
	JOBCLEANUPPOLICY_ON_SUCCESS JobCleanupPolicy = "on_success"
	JOBCLEANUPPOLICY_ALWAYS JobCleanupPolicy = "always"
	JOBCLEANUPPOLICY_NEVER JobCleanupPolicy = "never"
)