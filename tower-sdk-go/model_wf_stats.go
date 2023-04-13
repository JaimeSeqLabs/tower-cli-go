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
// WfStats struct for WfStats
type WfStats struct {
	ComputeTimeFmt string `json:"computeTimeFmt,omitempty"`
	CachedCount int32 `json:"cachedCount,omitempty"`
	FailedCount int32 `json:"failedCount,omitempty"`
	IgnoredCount int32 `json:"ignoredCount,omitempty"`
	SucceedCount int32 `json:"succeedCount,omitempty"`
	CachedCountFmt string `json:"cachedCountFmt,omitempty"`
	SucceedCountFmt string `json:"succeedCountFmt,omitempty"`
	FailedCountFmt string `json:"failedCountFmt,omitempty"`
	IgnoredCountFmt string `json:"ignoredCountFmt,omitempty"`
	CachedPct float32 `json:"cachedPct,omitempty"`
	FailedPct float32 `json:"failedPct,omitempty"`
	SucceedPct float32 `json:"succeedPct,omitempty"`
	IgnoredPct float32 `json:"ignoredPct,omitempty"`
	CachedDuration int64 `json:"cachedDuration,omitempty"`
	FailedDuration int64 `json:"failedDuration,omitempty"`
	SucceedDuration int64 `json:"succeedDuration,omitempty"`
}