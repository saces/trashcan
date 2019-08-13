# InlineResponse20033

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **string** | A token that can be used to paginate forwards with. | [optional] 
**Event** | [**map[string]interface{}**](map[string]interface{}.md) | Details of the requested event. | [optional] 
**EventsAfter** | [**[]map[string]interface{}**](map[string]interface{}.md) | A list of room events that happened just after the requested event, in chronological order. | [optional] 
**EventsBefore** | [**[]map[string]interface{}**](map[string]interface{}.md) | A list of room events that happened just before the requested event, in reverse-chronological order. | [optional] 
**Start** | **string** | A token that can be used to paginate backwards with. | [optional] 
**State** | [**[]map[string]interface{}**](map[string]interface{}.md) | The state of the room at the last event returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


