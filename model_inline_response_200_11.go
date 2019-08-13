/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse20011 struct {
	// The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Content map[string]interface{} `json:"content"`
	// The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	Type string `json:"type"`
}
