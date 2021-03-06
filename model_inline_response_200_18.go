/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse20018 struct {
	// For each key algorithm, the number of unclaimed one-time keys of that type currently held on the server for this device.
	OneTimeKeyCounts map[string]int32 `json:"one_time_key_counts"`
}
