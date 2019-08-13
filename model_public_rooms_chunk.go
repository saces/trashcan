/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PublicRoomsChunk struct {
	// Aliases of the room. May be empty.
	Aliases []string `json:"aliases,omitempty"`
	// The URL for the room's avatar, if one is set.
	AvatarUrl string `json:"avatar_url,omitempty"`
	// The canonical alias of the room, if any.
	CanonicalAlias string `json:"canonical_alias,omitempty"`
	// Whether guest users may join the room and participate in it. If they can, they will be subject to ordinary power level rules like any other user.
	GuestCanJoin bool `json:"guest_can_join"`
	// The name of the room, if any.
	Name string `json:"name,omitempty"`
	// The number of members joined to the room.
	NumJoinedMembers int32 `json:"num_joined_members"`
	// The ID of the room.
	RoomId string `json:"room_id"`
	// The topic of the room, if any.
	Topic string `json:"topic,omitempty"`
	// Whether the room may be viewed by guest users without joining.
	WorldReadable bool `json:"world_readable"`
}