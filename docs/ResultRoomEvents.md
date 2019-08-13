# ResultRoomEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | An approximate count of the total number of results found. | [optional] 
**Groups** | [**map[string]map[string]map[string]interface{}**](map.md) | Any groups that were requested.  The outer &#x60;&#x60;string&#x60;&#x60; key is the group key requested (eg: &#x60;&#x60;room_id&#x60;&#x60; or &#x60;&#x60;sender&#x60;&#x60;). The inner &#x60;&#x60;string&#x60;&#x60; key is the grouped value (eg: a room&#39;s ID or a user&#39;s ID). | [optional] 
**Highlights** | **[]string** | List of words which should be highlighted, useful for stemming which may change the query terms. | [optional] 
**NextBatch** | **string** | Token that can be used to get the next batch of results, by passing as the &#x60;next_batch&#x60; parameter to the next call. If this field is absent, there are no more results. | [optional] 
**Results** | [**[]Result**](Result.md) | List of results in the requested order. | [optional] 
**State** | [**map[string][]map[string]interface{}**](array.md) | The current state for every room in the results. This is included if the request had the &#x60;&#x60;include_state&#x60;&#x60; key set with a value of &#x60;&#x60;true&#x60;&#x60;.  The &#x60;&#x60;string&#x60;&#x60; key is the room ID for which the &#x60;&#x60;State Event&#x60;&#x60; array belongs to. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


