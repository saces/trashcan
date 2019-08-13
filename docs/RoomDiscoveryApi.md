# \RoomDiscoveryApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicRooms**](RoomDiscoveryApi.md#GetPublicRooms) | **Get** /_matrix/client/unstable/publicRooms | Lists the public rooms on the server.
[**QueryPublicRooms**](RoomDiscoveryApi.md#QueryPublicRooms) | **Post** /_matrix/client/unstable/publicRooms | Lists the public rooms on the server with optional filter.



## GetPublicRooms

> InlineResponse20026 GetPublicRooms(ctx, optional)
Lists the public rooms on the server.

Lists the public rooms on the server.  This API returns paginated responses. The rooms are ordered by the number of joined members, with the largest rooms first.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPublicRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPublicRoomsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of results returned. | 
 **since** | **optional.String**| A pagination token from a previous request, allowing clients to get the next (or previous) batch of rooms. The direction of pagination is specified solely by which token is supplied, rather than via an explicit flag. | 
 **server** | **optional.String**| The server to fetch the public room lists from. Defaults to the local server. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPublicRooms

> InlineResponse20027 QueryPublicRooms(ctx, body, optional)
Lists the public rooms on the server with optional filter.

Lists the public rooms on the server, with optional filter.  This API returns paginated responses. The rooms are ordered by the number of joined members, with the largest rooms first.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InlineObject18**](InlineObject18.md)|  | 
 **optional** | ***QueryPublicRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryPublicRoomsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | **optional.String**| The server to fetch the public room lists from. Defaults to the local server. | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

