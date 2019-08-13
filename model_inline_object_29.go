/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject29 struct {
	// The event ID the read marker should be located at. The event MUST belong to the room.
	MFullyRead string `json:"m.fully_read"`
	// The event ID to set the read receipt location at. This is equivalent to calling ``/receipt/m.read/$elsewhere:example.org`` and is provided here to save that extra call.
	MRead string `json:"m.read,omitempty"`
}
