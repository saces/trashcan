# \RoomUgpradesApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpgradeRoom**](RoomUgpradesApi.md#UpgradeRoom) | **Post** /_matrix/client/unstable/rooms/{roomId}/upgrade | Upgrades a room to a new room version.



## UpgradeRoom

> InlineResponse20038 UpgradeRoom(ctx, roomId, body)
Upgrades a room to a new room version.

Upgrades the given room to a particular room version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The ID of the room to upgrade. | 
**body** | [**InlineObject34**](InlineObject34.md)|  | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

