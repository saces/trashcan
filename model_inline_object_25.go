/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject25 struct {
	// The invitee's third party identifier.
	Address string `json:"address"`
	// The hostname+port of the identity server which should be used for third party identifier lookups.
	IdServer string `json:"id_server"`
	// The kind of address being passed in the address field, for example ``email``.
	Medium string `json:"medium"`
}
