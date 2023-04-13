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
// GoogleBatchConfig struct for GoogleBatchConfig
type GoogleBatchConfig struct {
	Location string `json:"location,omitempty"`
	WorkDir string `json:"workDir,omitempty"`
	Spot bool `json:"spot,omitempty"`
	BootDiskSizeGb int32 `json:"bootDiskSizeGb,omitempty"`
	CpuPlatform string `json:"cpuPlatform,omitempty"`
	MachineType string `json:"machineType,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	SshDaemon bool `json:"sshDaemon,omitempty"`
	SshImage string `json:"sshImage,omitempty"`
	DebugMode int32 `json:"debugMode,omitempty"`
	CopyImage string `json:"copyImage,omitempty"`
	UsePrivateAddress bool `json:"usePrivateAddress,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	PreRunScript string `json:"preRunScript,omitempty"`
	PostRunScript string `json:"postRunScript,omitempty"`
	HeadJobCpus int32 `json:"headJobCpus,omitempty"`
	HeadJobMemoryMb int32 `json:"headJobMemoryMb,omitempty"`
	NfsTarget string `json:"nfsTarget,omitempty"`
	NfsMount string `json:"nfsMount,omitempty"`
	Environment []ConfigEnvVariable `json:"environment,omitempty"`
	// property to select the compute config platform
	Discriminator string `json:"discriminator,omitempty"`
}