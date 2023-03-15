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
// TraceProgressDetail struct for TraceProgressDetail
type TraceProgressDetail struct {
	Index int32 `json:"index,omitempty"`
	Name string `json:"name,omitempty"`
	Pending int32 `json:"pending,omitempty"`
	Submitted int32 `json:"submitted,omitempty"`
	Running int32 `json:"running,omitempty"`
	Succeeded int32 `json:"succeeded,omitempty"`
	Cached int32 `json:"cached,omitempty"`
	Failed int32 `json:"failed,omitempty"`
	Aborted int32 `json:"aborted,omitempty"`
	Stored int32 `json:"stored,omitempty"`
	Ignored int32 `json:"ignored,omitempty"`
	Retries int32 `json:"retries,omitempty"`
	Terminated bool `json:"terminated,omitempty"`
	LoadCpus int64 `json:"loadCpus,omitempty"`
	LoadMemory int64 `json:"loadMemory,omitempty"`
	PeakRunning int32 `json:"peakRunning,omitempty"`
	PeakCpus int64 `json:"peakCpus,omitempty"`
	PeakMemory int64 `json:"peakMemory,omitempty"`
}
