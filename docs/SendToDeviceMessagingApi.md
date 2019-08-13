# \SendToDeviceMessagingApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendToDevice**](SendToDeviceMessagingApi.md#SendToDevice) | **Put** /_matrix/client/unstable/sendToDevice/{eventType}/{txnId} | Send an event to a given set of devices.



## SendToDevice

> SendToDevice(ctx, eventType, txnId, body)
Send an event to a given set of devices.

This endpoint is used to send send-to-device events to a set of client devices.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventType** | **string**| The type of event to send. | 
**txnId** | **string**| The transaction ID for this event. Clients should generate an ID unique across requests with the same access token; it will be used by the server to ensure idempotency of requests. | 
**body** | [**Body**](Body.md)|  | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

