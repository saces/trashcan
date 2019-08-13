# JoinedRoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | [**map[string]interface{}**](map[string]interface{}.md) | The private data that this user has attached to this room. | [optional] 
**Ephemeral** | [**map[string]interface{}**](map[string]interface{}.md) | The ephemeral events in the room that aren&#39;t recorded in the timeline or state of the room. e.g. typing. | [optional] 
**State** | [**map[string]interface{}**](map[string]interface{}.md) | Updates to the state, between the time indicated by the &#x60;&#x60;since&#x60;&#x60; parameter, and the start of the &#x60;&#x60;timeline&#x60;&#x60; (or all state up to the start of the &#x60;&#x60;timeline&#x60;&#x60;, if &#x60;&#x60;since&#x60;&#x60; is not given, or &#x60;&#x60;full_state&#x60;&#x60; is true).  N.B. state updates for &#x60;&#x60;m.room.member&#x60;&#x60; events will be incomplete if &#x60;&#x60;lazy_load_members&#x60;&#x60; is enabled in the &#x60;&#x60;/sync&#x60;&#x60; filter, and only return the member events required to display the senders of the timeline events in this response. | [optional] 
**Summary** | [**RoomSummary**](RoomSummary.md) |  | [optional] 
**Timeline** | [**map[string]interface{}**](map[string]interface{}.md) | The timeline of messages and state changes in the room. | [optional] 
**UnreadNotifications** | [**UnreadNotificationCounts**](Unread Notification Counts.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


