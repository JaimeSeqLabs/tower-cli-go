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
// SlurmComputeConfig struct for SlurmComputeConfig
type SlurmComputeConfig struct {
	WorkDir string `json:"workDir,omitempty"`
	LaunchDir string `json:"launchDir,omitempty"`
	UserName string `json:"userName,omitempty"`
	HostName string `json:"hostName,omitempty"`
	Port int32 `json:"port,omitempty"`
	HeadQueue string `json:"headQueue,omitempty"`
	ComputeQueue string `json:"computeQueue,omitempty"`
	MaxQueueSize int32 `json:"maxQueueSize,omitempty"`
	HeadJobOptions string `json:"headJobOptions,omitempty"`
	PropagateHeadJobOptions bool `json:"propagateHeadJobOptions,omitempty"`
	PreRunScript string `json:"preRunScript,omitempty"`
	PostRunScript string `json:"postRunScript,omitempty"`
	Environment []ConfigEnvVariable `json:"environment,omitempty"`
	// property to select the compute config platform
	Discriminator string `json:"discriminator,omitempty"`
}
