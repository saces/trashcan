/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Filters to be applied to room data.
type RoomFilter struct {
	// The per user account data to include for rooms.
	AccountData map[string]interface{} `json:"account_data,omitempty"`
	// The events that aren't recorded in the room history, e.g. typing and receipts, to include for rooms.
	Ephemeral map[string]interface{} `json:"ephemeral,omitempty"`
	// Include rooms that the user has left in the sync, default false
	IncludeLeave bool `json:"include_leave,omitempty"`
	// A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the ``'rooms'`` filter. This filter is applied before the filters in ``ephemeral``, ``state``, ``timeline`` or ``account_data``
	NotRooms []string `json:"not_rooms,omitempty"`
	// A list of room IDs to include. If this list is absent then all rooms are included. This filter is applied before the filters in ``ephemeral``, ``state``, ``timeline`` or ``account_data``
	Rooms []string `json:"rooms,omitempty"`
	// The state events to include for rooms.
	State map[string]interface{} `json:"state,omitempty"`
	// The message and state update events to include for rooms.
	Timeline map[string]interface{} `json:"timeline,omitempty"`
}
