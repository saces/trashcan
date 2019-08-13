# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Additional authentication information for the user-interactive authentication API. | [optional] 
**IdServer** | **string** | The identity server to unbind all of the user&#39;s 3PIDs from. If not provided, the homeserver MUST use the &#x60;&#x60;id_server&#x60;&#x60; that was originally use to bind each identifier. If the homeserver does not know which &#x60;&#x60;id_server&#x60;&#x60; that was, it must return an &#x60;&#x60;id_server_unbind_result&#x60;&#x60; of &#x60;&#x60;no-support&#x60;&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


