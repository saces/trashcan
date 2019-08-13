# InlineResponse20026

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunk** | [**[]PublicRoomsChunk**](PublicRoomsChunk.md) | A paginated chunk of public rooms. | 
**NextBatch** | **string** | A pagination token for the response. The absence of this token means there are no more results to fetch and the client should stop paginating. | [optional] 
**PrevBatch** | **string** | A pagination token that allows fetching previous results. The absence of this token means there are no results before this batch, i.e. this is the first batch. | [optional] 
**TotalRoomCountEstimate** | **int32** | An estimate on the total number of public rooms, if the server has an estimate. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


