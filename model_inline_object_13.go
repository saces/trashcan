/*
 * Matrix Client-Server API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: unstable
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject13 struct {
	// Identity keys for the device. May be absent if no new identity keys are required.
	DeviceKeys map[string]interface{} `json:"device_keys,omitempty"`
	// One-time public keys for \"pre-key\" messages.  The names of the properties should be in the format ``<algorithm>:<key_id>``. The format of the key is determined by the `key algorithm <#key-algorithms>`_.  May be absent if no new one-time keys are required.
	OneTimeKeys map[string]interface{} `json:"one_time_keys,omitempty"`
}