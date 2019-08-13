/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse20015 struct {
	// The Matrix User IDs of all users who updated their device identity keys.
	Changed []string `json:"changed,omitempty"`
	// The Matrix User IDs of all users who may have left all the end-to-end encrypted rooms they previously shared with the user.
	Left []string `json:"left,omitempty"`
}
