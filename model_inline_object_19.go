/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject19 struct {
	// A string that will allow the user to identify what application owns this pusher.
	AppDisplayName string `json:"app_display_name"`
	// This is a reverse-DNS style identifier for the application. It is recommended that this end with the platform, such that different platform versions get different app identifiers. Max length, 64 chars.  If the ``kind`` is ``\"email\"``, this is ``\"m.email\"``.
	AppId string `json:"app_id"`
	// If true, the homeserver should add another pusher with the given pushkey and App ID in addition to any others with different user IDs. Otherwise, the homeserver must remove any other pushers with the same App ID and pushkey for different users. The default is ``false``.
	Append bool `json:"append,omitempty"`
	Data PusherData1 `json:"data"`
	// A string that will allow the user to identify what device owns this pusher.
	DeviceDisplayName string `json:"device_display_name"`
	// The kind of pusher to configure. ``\"http\"`` makes a pusher that sends HTTP pokes. ``\"email\"`` makes a pusher that emails the user with unread notifications. ``null`` deletes the pusher.
	Kind string `json:"kind"`
	// The preferred language for receiving notifications (e.g. 'en' or 'en-US').
	Lang string `json:"lang"`
	// This string determines which set of device specific rules this pusher executes.
	ProfileTag string `json:"profile_tag,omitempty"`
	// This is a unique identifier for this pusher. The value you should use for this is the routing or destination address information for the notification, for example, the APNS token for APNS or the Registration ID for GCM. If your notification client has no such concept, use any unique identifier. Max length, 512 bytes.  If the ``kind`` is ``\"email\"``, this is the email address to send notifications to.
	Pushkey string `json:"pushkey"`
}
