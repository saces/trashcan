# EventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **string** | Pagination token for the end of the chunk | [optional] 
**EventsAfter** | [**[]RoomEvent**](RoomEvent.md) | Events just after the result. | [optional] 
**EventsBefore** | [**[]RoomEvent**](RoomEvent.md) | Events just before the result. | [optional] 
**ProfileInfo** | [**map[string]UserProfile**](User Profile.md) | The historic profile information of the users that sent the events returned.  The &#x60;&#x60;string&#x60;&#x60; key is the user ID for which the profile belongs to. | [optional] 
**Start** | **string** | Pagination token for the start of the chunk | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


