/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse20045 struct {
	// The password to use.
	Password string `json:"password"`
	// The time-to-live in seconds
	Ttl int32 `json:"ttl"`
	// A list of TURN URIs
	Uris []string `json:"uris"`
	// The username to use.
	Username string `json:"username"`
}