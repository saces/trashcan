# InlineResponse20016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failures** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | If any remote homeservers could not be reached, they are recorded here. The names of the properties are the names of the unreachable servers.  If the homeserver could be reached, but the user or device was unknown, no failure is recorded. Instead, the corresponding user or device is missing from the &#x60;&#x60;one_time_keys&#x60;&#x60; result. | [optional] 
**OneTimeKeys** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | One-time keys for the queried devices. A map from user ID, to a map from devices to a map from &#x60;&#x60;&lt;algorithm&gt;:&lt;key_id&gt;&#x60;&#x60; to the key object.  See the &#x60;key algorithms &lt;#key-algorithms&gt;&#x60;_ section for information on the Key Object format. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


