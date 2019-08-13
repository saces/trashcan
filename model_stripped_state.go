/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A stripped down state event, with only the ``type``, ``state_key``, ``sender``, and ``content`` keys.
type StrippedState struct {
	// The ``content`` for the event.
	Content map[string]interface{} `json:"content"`
	// The ``sender`` for the event.
	Sender string `json:"sender"`
	// The ``state_key`` for the event.
	StateKey string `json:"state_key"`
	// The ``type`` for the event.
	Type string `json:"type"`
}
