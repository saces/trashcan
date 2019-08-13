/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A dictionary of information for the pusher implementation itself. If ``kind`` is ``http``, this should contain ``url`` which is the URL to use to send notifications to.
type PusherData1 struct {
	// The format to send notifications in to Push Gateways if the ``kind`` is ``http``. The details about what fields the homeserver should send to the push gateway are defined in the `Push Gateway Specification`_. Currently the only format available is 'event_id_only'.
	Format string `json:"format,omitempty"`
	// Required if ``kind`` is ``http``. The URL to use to send notifications to. MUST be an HTTPS URL with a path of  ``/_matrix/push/v1/notify``.
	Url string `json:"url,omitempty"`
}
