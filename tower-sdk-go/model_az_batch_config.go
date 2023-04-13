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
// AzBatchConfig struct for AzBatchConfig
type AzBatchConfig struct {
	WorkDir string `json:"workDir,omitempty"`
	PreRunScript string `json:"preRunScript,omitempty"`
	PostRunScript string `json:"postRunScript,omitempty"`
	Region string `json:"region,omitempty"`
	HeadPool string `json:"headPool,omitempty"`
	AutoPoolMode bool `json:"autoPoolMode,omitempty"`
	Forge AzBatchForgeConfig `json:"forge,omitempty"`
	TokenDuration string `json:"tokenDuration,omitempty"`
	DeleteJobsOnCompletion JobCleanupPolicy `json:"deleteJobsOnCompletion,omitempty"`
	DeletePoolsOnCompletion bool `json:"deletePoolsOnCompletion,omitempty"`
	Environment []ConfigEnvVariable `json:"environment,omitempty"`
	// property to select the compute config platform
	Discriminator string `json:"discriminator,omitempty"`
}