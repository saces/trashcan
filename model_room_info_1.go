/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RoomInfo1 struct {
	// The private data that this user has attached to this room.
	AccountData []map[string]interface{} `json:"account_data,omitempty"`
	// The user's membership state in this room.
	Membership string `json:"membership,omitempty"`
	Messages PaginationChunk `json:"messages,omitempty"`
	// The ID of this room.
	RoomId string `json:"room_id"`
	// If the user is a member of the room this will be the current state of the room as a list of events. If the user has left the room this will be the state of the room when they left it.
	State []map[string]interface{} `json:"state,omitempty"`
	// Whether this room is visible to the ``/publicRooms`` API or not.\"
	Visibility string `json:"visibility,omitempty"`
}
