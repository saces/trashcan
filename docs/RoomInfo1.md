# RoomInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | [**[]map[string]interface{}**](map[string]interface{}.md) | The private data that this user has attached to this room. | [optional] 
**Membership** | **string** | The user&#39;s membership state in this room. | [optional] 
**Messages** | [**PaginationChunk**](PaginationChunk.md) |  | [optional] 
**RoomId** | **string** | The ID of this room. | 
**State** | [**[]map[string]interface{}**](map[string]interface{}.md) | If the user is a member of the room this will be the current state of the room as a list of events. If the user has left the room this will be the state of the room when they left it. | [optional] 
**Visibility** | **string** | Whether this room is visible to the &#x60;&#x60;/publicRooms&#x60;&#x60; API or not.\&quot; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


