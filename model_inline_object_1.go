/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject1 struct {
	// The third party address being removed.
	Address string `json:"address"`
	// The identity server to unbind from. If not provided, the homeserver MUST use the ``id_server`` the identifier was added through. If the homeserver does not know the original ``id_server``, it MUST return a ``id_server_unbind_result`` of ``no-support``.
	IdServer string `json:"id_server,omitempty"`
	// The medium of the third party identifier being removed.
	Medium string `json:"medium"`
}