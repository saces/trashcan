/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type JoinedRoom struct {
	// The private data that this user has attached to this room.
	AccountData map[string]interface{} `json:"account_data,omitempty"`
	// The ephemeral events in the room that aren't recorded in the timeline or state of the room. e.g. typing.
	Ephemeral map[string]interface{} `json:"ephemeral,omitempty"`
	// Updates to the state, between the time indicated by the ``since`` parameter, and the start of the ``timeline`` (or all state up to the start of the ``timeline``, if ``since`` is not given, or ``full_state`` is true).  N.B. state updates for ``m.room.member`` events will be incomplete if ``lazy_load_members`` is enabled in the ``/sync`` filter, and only return the member events required to display the senders of the timeline events in this response.
	State map[string]interface{} `json:"state,omitempty"`
	Summary RoomSummary `json:"summary,omitempty"`
	// The timeline of messages and state changes in the room.
	Timeline map[string]interface{} `json:"timeline,omitempty"`
	UnreadNotifications UnreadNotificationCounts `json:"unread_notifications,omitempty"`
}
