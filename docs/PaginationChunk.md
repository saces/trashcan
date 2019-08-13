# PaginationChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunk** | [**[]map[string]interface{}**](map[string]interface{}.md) | If the user is a member of the room this will be a list of the most recent messages for this room. If the user has left the room this will be the messages that preceeded them leaving. This array will consist of at most &#x60;&#x60;limit&#x60;&#x60; elements. | 
**End** | **string** | A token which correlates to the last value in &#x60;&#x60;chunk&#x60;&#x60;. Used for pagination. | 
**Start** | **string** | A token which correlates to the first value in &#x60;&#x60;chunk&#x60;&#x60;. Used for pagination. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


