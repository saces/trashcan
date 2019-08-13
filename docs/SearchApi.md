# \SearchApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchApi.md#Search) | **Post** /_matrix/client/unstable/search | Perform a server-side search.



## Search

> Results Search(ctx, optional)
Perform a server-side search.

Performs a full text search across different categories.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextBatch** | **optional.String**| The point to return events from. If given, this should be a &#x60;&#x60;next_batch&#x60;&#x60; result from a previous call to this endpoint. | 
 **body** | [**optional.Interface of InlineObject35**](InlineObject35.md)|  | 

### Return type

[**Results**](Results.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

