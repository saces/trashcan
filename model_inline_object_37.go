/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject37 struct {
	// The maximum number of results to return. Defaults to 10.
	Limit int32 `json:"limit,omitempty"`
	// The term to search for
	SearchTerm string `json:"search_term"`
}