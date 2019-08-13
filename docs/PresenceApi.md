# \PresenceApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPresence**](PresenceApi.md#GetPresence) | **Get** /_matrix/client/unstable/presence/{userId}/status | Get this user&#39;s presence state.
[**SetPresence**](PresenceApi.md#SetPresence) | **Put** /_matrix/client/unstable/presence/{userId}/status | Update this user&#39;s presence state.



## GetPresence

> InlineResponse20022 GetPresence(ctx, userId)
Get this user's presence state.

Get the given user's presence state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose presence state to get. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPresence

> map[string]interface{} SetPresence(ctx, userId, presenceState)
Update this user's presence state.

This API sets the given user's presence state. When setting the status, the activity time is updated to reflect that activity; the client does not need to specify the ``last_active_ago`` field. You cannot set the presence state of another user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose presence state to update. | 
**presenceState** | [**InlineObject15**](InlineObject15.md)|  | 

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

