/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ThirdPartyIdentifier struct {
	// The timestamp, in milliseconds, when the homeserver associated the third party identifier with the user.
	AddedAt int64 `json:"added_at"`
	// The third party identifier address.
	Address string `json:"address"`
	// The medium of the third party identifier.
	Medium string `json:"medium"`
	// The timestamp, in milliseconds, when the identifier was validated by the identity server.
	ValidatedAt int64 `json:"validated_at"`
}