# RoomFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | [**map[string]interface{}**](map[string]interface{}.md) | The per user account data to include for rooms. | [optional] 
**Ephemeral** | [**map[string]interface{}**](map[string]interface{}.md) | The events that aren&#39;t recorded in the room history, e.g. typing and receipts, to include for rooms. | [optional] 
**IncludeLeave** | **bool** | Include rooms that the user has left in the sync, default false | [optional] 
**NotRooms** | **[]string** | A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the &#x60;&#x60;&#39;rooms&#39;&#x60;&#x60; filter. This filter is applied before the filters in &#x60;&#x60;ephemeral&#x60;&#x60;, &#x60;&#x60;state&#x60;&#x60;, &#x60;&#x60;timeline&#x60;&#x60; or &#x60;&#x60;account_data&#x60;&#x60; | [optional] 
**Rooms** | **[]string** | A list of room IDs to include. If this list is absent then all rooms are included. This filter is applied before the filters in &#x60;&#x60;ephemeral&#x60;&#x60;, &#x60;&#x60;state&#x60;&#x60;, &#x60;&#x60;timeline&#x60;&#x60; or &#x60;&#x60;account_data&#x60;&#x60; | [optional] 
**State** | [**map[string]interface{}**](map[string]interface{}.md) | The state events to include for rooms. | [optional] 
**Timeline** | [**map[string]interface{}**](map[string]interface{}.md) | The message and state update events to include for rooms. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


