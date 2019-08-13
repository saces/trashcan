# \ApplicationServiceRoomDirectoryManagementApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateAppserviceRoomDirectoryVsibility**](ApplicationServiceRoomDirectoryManagementApi.md#UpdateAppserviceRoomDirectoryVsibility) | **Put** /_matrix/client/unstable/directory/list/appservice/{networkId}/{roomId} | Updates a room&#39;s visibility in the application service&#39;s room directory.



## UpdateAppserviceRoomDirectoryVsibility

> map[string]interface{} UpdateAppserviceRoomDirectoryVsibility(ctx, networkId, roomId, body)
Updates a room's visibility in the application service's room directory.

Updates the visibility of a given room on the application service's room directory.  This API is similar to the room directory visibility API used by clients to update the homeserver's more general room directory.  This API requires the use of an application service access token (``as_token``) instead of a typical client's access_token. This API cannot be invoked by users who are not identified as application services.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**| The protocol (network) ID to update the room list for. This would have been provided by the application service as being listed as a supported protocol. | 
**roomId** | **string**| The room ID to add to the directory. | 
**body** | [**InlineObject8**](InlineObject8.md)|  | 

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

