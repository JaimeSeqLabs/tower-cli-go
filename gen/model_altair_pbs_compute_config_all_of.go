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
// AltairPbsComputeConfigAllOf struct for AltairPbsComputeConfigAllOf
type AltairPbsComputeConfigAllOf struct {
	Environment []ConfigEnvVariable `json:"environment,omitempty"`
	// property to select the compute config platform
	Discriminator string `json:"discriminator,omitempty"`
}
