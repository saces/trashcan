# \RoomDirectoryApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoomAlias**](RoomDirectoryApi.md#DeleteRoomAlias) | **Delete** /_matrix/client/unstable/directory/room/{roomAlias} | Remove a mapping of room alias to room ID.
[**GetRoomIdByAlias**](RoomDirectoryApi.md#GetRoomIdByAlias) | **Get** /_matrix/client/unstable/directory/room/{roomAlias} | Get the room ID corresponding to this room alias.
[**SetRoomAlias**](RoomDirectoryApi.md#SetRoomAlias) | **Put** /_matrix/client/unstable/directory/room/{roomAlias} | Create a new mapping from room alias to room ID.



## DeleteRoomAlias

> map[string]interface{} DeleteRoomAlias(ctx, roomAlias)
Remove a mapping of room alias to room ID.

Remove a mapping of room alias to room ID.  Servers may choose to implement additional access control checks here, for instance that room aliases can only be deleted by their creator or a server administrator.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomAlias** | **string**| The room alias to remove. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomIdByAlias

> InlineResponse2009 GetRoomIdByAlias(ctx, roomAlias)
Get the room ID corresponding to this room alias.

Requests that the server resolve a room alias to a room ID.  The server will use the federation API to resolve the alias if the domain part of the alias does not correspond to the server's own domain.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomAlias** | **string**| The room alias. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoomAlias

> map[string]interface{} SetRoomAlias(ctx, roomAlias, body)
Create a new mapping from room alias to room ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomAlias** | **string**| The room alias to set. | 
**body** | [**InlineObject9**](InlineObject9.md)|  | 

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

