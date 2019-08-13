# \ReportingContentApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportContent**](ReportingContentApi.md#ReportContent) | **Post** /_matrix/client/unstable/rooms/{roomId}/report/{eventId} | Reports an event as inappropriate.



## ReportContent

> map[string]interface{} ReportContent(ctx, roomId, eventId, optional)
Reports an event as inappropriate.

Reports an event as inappropriate to the server, which may then notify the appropriate people.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room in which the event being reported is located. | 
**eventId** | **string**| The event to report. | 
 **optional** | ***ReportContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InlineObject31**](InlineObject31.md)|  | 

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

