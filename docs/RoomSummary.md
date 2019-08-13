# RoomSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MHeroes** | **[]string** | The users which can be used to generate a room name if the room does not have one. Required if the room&#39;s &#x60;&#x60;m.room.name&#x60;&#x60; or &#x60;&#x60;m.room.canonical_alias&#x60;&#x60; state events are unset or empty.  This should be the first 5 members of the room, ordered by stream ordering, which are joined or invited. The list must never include the client&#39;s own user ID. When no joined or invited members are available, this should consist of the banned and left users. More than 5 members may be provided, however less than 5 should only be provided when there are less than 5 members to represent.  When lazy-loading room members is enabled, the membership events for the heroes MUST be included in the &#x60;&#x60;state&#x60;&#x60;, unless they are redundant. When the list of users changes, the server notifies the client by sending a fresh list of heroes. If there are no changes since the last sync, this field may be omitted. | [optional] 
**MInvitedMemberCount** | **int32** | The number of users with &#x60;&#x60;membership&#x60;&#x60; of &#x60;&#x60;invite&#x60;&#x60;. If this field has not changed since the last sync, it may be omitted. Required otherwise. | [optional] 
**MJoinedMemberCount** | **int32** | The number of users with &#x60;&#x60;membership&#x60;&#x60; of &#x60;&#x60;join&#x60;&#x60;, including the client&#39;s own user ID. If this field has not changed since the last sync, it may be omitted. Required otherwise. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


