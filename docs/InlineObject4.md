# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationContent** | [**map[string]interface{}**](.md) | Extra keys, such as &#x60;&#x60;m.federate&#x60;&#x60;, to be added to the content of the &#x60;m.room.create&#x60;_ event. The server will clobber the following keys: &#x60;&#x60;creator&#x60;&#x60;, &#x60;&#x60;room_version&#x60;&#x60;. Future versions of the specification may allow the server to clobber other keys. | [optional] 
**InitialState** | [**[]StateEvent**](StateEvent.md) | A list of state events to set in the new room. This allows the user to override the default state events set in the new room. The expected format of the state events are an object with type, state_key and content keys set.  Takes precedence over events set by &#x60;&#x60;preset&#x60;&#x60;, but gets overriden by &#x60;&#x60;name&#x60;&#x60; and &#x60;&#x60;topic&#x60;&#x60; keys. | [optional] 
**Invite** | **[]string** | A list of user IDs to invite to the room. This will tell the server to invite everyone in the list to the newly created room. | [optional] 
**Invite3pid** | [**[]Invite3pid**](Invite3pid.md) | A list of objects representing third party IDs to invite into the room. | [optional] 
**IsDirect** | **bool** | This flag makes the server set the &#x60;&#x60;is_direct&#x60;&#x60; flag on the &#x60;&#x60;m.room.member&#x60;&#x60; events sent to the users in &#x60;&#x60;invite&#x60;&#x60; and &#x60;&#x60;invite_3pid&#x60;&#x60;. See &#x60;Direct Messaging&#x60;_ for more information. | [optional] 
**Name** | **string** | If this is included, an &#x60;&#x60;m.room.name&#x60;&#x60; event will be sent into the room to indicate the name of the room. See Room Events for more information on &#x60;&#x60;m.room.name&#x60;&#x60;. | [optional] 
**PowerLevelContentOverride** | [**map[string]interface{}**](.md) | The power level content to override in the default power level event. This object is applied on top of the generated &#x60;m.room.power_levels&#x60;_ event content prior to it being sent to the room. Defaults to overriding nothing. | [optional] 
**Preset** | **string** | Convenience parameter for setting various default state events based on a preset.  If unspecified, the server should use the &#x60;&#x60;visibility&#x60;&#x60; to determine which preset to use. A visbility of &#x60;&#x60;public&#x60;&#x60; equates to a preset of &#x60;&#x60;public_chat&#x60;&#x60; and &#x60;&#x60;private&#x60;&#x60; visibility equates to a preset of &#x60;&#x60;private_chat&#x60;&#x60;. | [optional] 
**RoomAliasName** | **string** | The desired room alias **local part**. If this is included, a room alias will be created and mapped to the newly created room. The alias will belong on the *same* homeserver which created the room. For example, if this was set to \&quot;foo\&quot; and sent to the homeserver \&quot;example.com\&quot; the complete room alias would be &#x60;&#x60;#foo:example.com&#x60;&#x60;.  The complete room alias will become the canonical alias for the room. | [optional] 
**RoomVersion** | **string** | The room version to set for the room. If not provided, the homeserver is to use its configured default. If provided, the homeserver will return a 400 error with the errcode &#x60;&#x60;M_UNSUPPORTED_ROOM_VERSION&#x60;&#x60; if it does not support the room version. | [optional] 
**Topic** | **string** | If this is included, an &#x60;&#x60;m.room.topic&#x60;&#x60; event will be sent into the room to indicate the topic for the room. See Room Events for more information on &#x60;&#x60;m.room.topic&#x60;&#x60;. | [optional] 
**Visibility** | **string** | A &#x60;&#x60;public&#x60;&#x60; visibility indicates that the room will be shown in the published room list. A &#x60;&#x60;private&#x60;&#x60; visibility will hide the room from the published room list. Rooms default to &#x60;&#x60;private&#x60;&#x60; visibility if this key is not included. NB: This should not be confused with &#x60;&#x60;join_rules&#x60;&#x60; which also uses the word &#x60;&#x60;public&#x60;&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


