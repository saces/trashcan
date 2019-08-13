/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The events and state surrounding the requested event.
type InlineResponse20033 struct {
	// A token that can be used to paginate forwards with.
	End string `json:"end,omitempty"`
	// Details of the requested event.
	Event map[string]interface{} `json:"event,omitempty"`
	// A list of room events that happened just after the requested event, in chronological order.
	EventsAfter []map[string]interface{} `json:"events_after,omitempty"`
	// A list of room events that happened just before the requested event, in reverse-chronological order.
	EventsBefore []map[string]interface{} `json:"events_before,omitempty"`
	// A token that can be used to paginate backwards with.
	Start string `json:"start,omitempty"`
	// The state of the room at the last event returned.
	State []map[string]interface{} `json:"state,omitempty"`
}
