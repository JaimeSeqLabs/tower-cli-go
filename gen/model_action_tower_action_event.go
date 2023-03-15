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
import (
	"time"
)
// ActionTowerActionEvent struct for ActionTowerActionEvent
type ActionTowerActionEvent struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	WorkflowId string `json:"workflowId,omitempty"`
	Discriminator string `json:"discriminator,omitempty"`
}
