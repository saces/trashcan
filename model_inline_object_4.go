/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject4 struct {
	// Extra keys, such as ``m.federate``, to be added to the content of the `m.room.create`_ event. The server will clobber the following keys: ``creator``, ``room_version``. Future versions of the specification may allow the server to clobber other keys.
	CreationContent map[string]interface{} `json:"creation_content,omitempty"`
	// A list of state events to set in the new room. This allows the user to override the default state events set in the new room. The expected format of the state events are an object with type, state_key and content keys set.  Takes precedence over events set by ``preset``, but gets overriden by ``name`` and ``topic`` keys.
	InitialState []StateEvent `json:"initial_state,omitempty"`
	// A list of user IDs to invite to the room. This will tell the server to invite everyone in the list to the newly created room.
	Invite []string `json:"invite,omitempty"`
	// A list of objects representing third party IDs to invite into the room.
	Invite3pid []Invite3pid `json:"invite_3pid,omitempty"`
	// This flag makes the server set the ``is_direct`` flag on the ``m.room.member`` events sent to the users in ``invite`` and ``invite_3pid``. See `Direct Messaging`_ for more information.
	IsDirect bool `json:"is_direct,omitempty"`
	// If this is included, an ``m.room.name`` event will be sent into the room to indicate the name of the room. See Room Events for more information on ``m.room.name``.
	Name string `json:"name,omitempty"`
	// The power level content to override in the default power level event. This object is applied on top of the generated `m.room.power_levels`_ event content prior to it being sent to the room. Defaults to overriding nothing.
	PowerLevelContentOverride map[string]interface{} `json:"power_level_content_override,omitempty"`
	// Convenience parameter for setting various default state events based on a preset.  If unspecified, the server should use the ``visibility`` to determine which preset to use. A visbility of ``public`` equates to a preset of ``public_chat`` and ``private`` visibility equates to a preset of ``private_chat``.
	Preset string `json:"preset,omitempty"`
	// The desired room alias **local part**. If this is included, a room alias will be created and mapped to the newly created room. The alias will belong on the *same* homeserver which created the room. For example, if this was set to \"foo\" and sent to the homeserver \"example.com\" the complete room alias would be ``#foo:example.com``.  The complete room alias will become the canonical alias for the room.
	RoomAliasName string `json:"room_alias_name,omitempty"`
	// The room version to set for the room. If not provided, the homeserver is to use its configured default. If provided, the homeserver will return a 400 error with the errcode ``M_UNSUPPORTED_ROOM_VERSION`` if it does not support the room version.
	RoomVersion string `json:"room_version,omitempty"`
	// If this is included, an ``m.room.topic`` event will be sent into the room to indicate the topic for the room. See Room Events for more information on ``m.room.topic``.
	Topic string `json:"topic,omitempty"`
	// A ``public`` visibility indicates that the room will be shown in the published room list. A ``private`` visibility will hide the room from the published room list. Rooms default to ``private`` visibility if this key is not included. NB: This should not be confused with ``join_rules`` which also uses the word ``public``.
	Visibility string `json:"visibility,omitempty"`
}