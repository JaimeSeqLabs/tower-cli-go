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
// ActionResponseDto struct for ActionResponseDto
type ActionResponseDto struct {
	Id string `json:"id,omitempty"`
	Launch Launch `json:"launch,omitempty"`
	Name string `json:"name,omitempty"`
	HookId string `json:"hookId,omitempty"`
	HookUrl string `json:"hookUrl,omitempty"`
	Message string `json:"message,omitempty"`
	Source ActionSource `json:"source,omitempty"`
	Status ActionStatus `json:"status,omitempty"`
	Config ActionConfigType `json:"config,omitempty"`
	Event ActionEventType `json:"event,omitempty"`
	LastSeen time.Time `json:"lastSeen,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Labels []LabelDbDto `json:"labels,omitempty"`
}
