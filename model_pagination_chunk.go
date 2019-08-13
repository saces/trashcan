/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The pagination chunk for this room.
type PaginationChunk struct {
	// If the user is a member of the room this will be a list of the most recent messages for this room. If the user has left the room this will be the messages that preceeded them leaving. This array will consist of at most ``limit`` elements.
	Chunk []map[string]interface{} `json:"chunk"`
	// A token which correlates to the last value in ``chunk``. Used for pagination.
	End string `json:"end"`
	// A token which correlates to the first value in ``chunk``. Used for pagination.
	Start string `json:"start"`
}
