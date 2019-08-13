# InlineResponse20041

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | [**map[string]interface{}**](map[string]interface{}.md) | The user account data that isn&#39;t associated with rooms to include. | [optional] 
**EventFields** | **[]string** | List of event fields to include. If this list is absent then all fields are included. The entries may include &#39;.&#39; characters to indicate sub-fields. So [&#39;content.body&#39;] will include the &#39;body&#39; field of the &#39;content&#39; object. A literal &#39;.&#39; character in a field name may be escaped using a &#39;\\\\&#39;. A server may include more fields than were requested. | [optional] 
**EventFormat** | **string** | The format to use for events. &#39;client&#39; will return the events in a format suitable for clients. &#39;federation&#39; will return the raw event as received over federation. The default is &#39;client&#39;. | [optional] 
**Presence** | [**map[string]interface{}**](map[string]interface{}.md) | The presence updates to include. | [optional] 
**Room** | [**RoomFilter**](RoomFilter.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


