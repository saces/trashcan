# InlineResponse20036

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunk** | [**[]RoomEvent**](RoomEvent.md) | A list of room events. The order depends on the &#x60;&#x60;dir&#x60;&#x60; parameter. For &#x60;&#x60;dir&#x3D;b&#x60;&#x60; events will be in reverse-chronological order, for &#x60;&#x60;dir&#x3D;f&#x60;&#x60; in chronological order, so that events start at the &#x60;&#x60;from&#x60;&#x60; point. | [optional] 
**End** | **string** | The token the pagination ends at. If &#x60;&#x60;dir&#x3D;b&#x60;&#x60; this token should be used again to request even earlier events. | [optional] 
**Start** | **string** | The token the pagination starts from. If &#x60;&#x60;dir&#x3D;b&#x60;&#x60; this will be the token supplied in &#x60;&#x60;from&#x60;&#x60;. | [optional] 
**State** | [**[]map[string]interface{}**](map[string]interface{}.md) | A list of state events relevant to showing the &#x60;&#x60;chunk&#x60;&#x60;. For example, if &#x60;&#x60;lazy_load_members&#x60;&#x60; is enabled in the filter then this may contain the membership events for the senders of events in the &#x60;&#x60;chunk&#x60;&#x60;.  Unless &#x60;&#x60;include_redundant_members&#x60;&#x60; is &#x60;&#x60;true&#x60;&#x60;, the server may remove membership events which would have already been sent to the client in prior calls to this endpoint, assuming the membership of those members has not changed. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


