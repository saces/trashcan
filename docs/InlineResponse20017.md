# InlineResponse20017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceKeys** | [**map[string]map[string]map[string]interface{}**](map.md) | Information on the queried devices. A map from user ID, to a map from device ID to device information.  For each device, the information returned will be the same as uploaded via &#x60;&#x60;/keys/upload&#x60;&#x60;, with the addition of an &#x60;&#x60;unsigned&#x60;&#x60; property. | [optional] 
**Failures** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | If any remote homeservers could not be reached, they are recorded here. The names of the properties are the names of the unreachable servers.  If the homeserver could be reached, but the user or device was unknown, no failure is recorded. Instead, the corresponding user or device is missing from the &#x60;&#x60;device_keys&#x60;&#x60; result. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


