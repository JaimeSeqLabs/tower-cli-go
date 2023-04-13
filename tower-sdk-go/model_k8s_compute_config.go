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
// K8sComputeConfig struct for K8sComputeConfig
type K8sComputeConfig struct {
	WorkDir string `json:"workDir,omitempty"`
	PreRunScript string `json:"preRunScript,omitempty"`
	PostRunScript string `json:"postRunScript,omitempty"`
	Server string `json:"server,omitempty"`
	SslCert string `json:"sslCert,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ComputeServiceAccount string `json:"computeServiceAccount,omitempty"`
	HeadServiceAccount string `json:"headServiceAccount,omitempty"`
	StorageClaimName string `json:"storageClaimName,omitempty"`
	StorageMountPath string `json:"storageMountPath,omitempty"`
	PodCleanup PodCleanupPolicy `json:"podCleanup,omitempty"`
	HeadPodSpec string `json:"headPodSpec,omitempty"`
	ServicePodSpec string `json:"servicePodSpec,omitempty"`
	Environment []ConfigEnvVariable `json:"environment,omitempty"`
	HeadJobCpus int32 `json:"headJobCpus,omitempty"`
	HeadJobMemoryMb int32 `json:"headJobMemoryMb,omitempty"`
	// property to select the compute config platform
	Discriminator string `json:"discriminator,omitempty"`
}