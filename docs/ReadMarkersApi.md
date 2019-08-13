# \ReadMarkersApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetReadMarker**](ReadMarkersApi.md#SetReadMarker) | **Post** /_matrix/client/unstable/rooms/{roomId}/read_markers | Set the position of the read marker for a room.



## SetReadMarker

> map[string]interface{} SetReadMarker(ctx, roomId, body)
Set the position of the read marker for a room.

Sets the position of the read marker for a given room, and optionally the read receipt's location.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room ID to set the read marker in for the user. | 
**body** | [**InlineObject29**](InlineObject29.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

