# \EndToEndEncryptionApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimKeys**](EndToEndEncryptionApi.md#ClaimKeys) | **Post** /_matrix/client/unstable/keys/claim | Claim one-time encryption keys.
[**GetKeysChanges**](EndToEndEncryptionApi.md#GetKeysChanges) | **Get** /_matrix/client/unstable/keys/changes | Query users with recent device key updates.
[**QueryKeys**](EndToEndEncryptionApi.md#QueryKeys) | **Post** /_matrix/client/unstable/keys/query | Download device identity keys.
[**UploadKeys**](EndToEndEncryptionApi.md#UploadKeys) | **Post** /_matrix/client/unstable/keys/upload | Upload end-to-end encryption keys.



## ClaimKeys

> InlineResponse20016 ClaimKeys(ctx, optional)
Claim one-time encryption keys.

Claims one-time keys for use in pre-key messages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClaimKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClaimKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**optional.Interface of InlineObject11**](InlineObject11.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeysChanges

> InlineResponse20015 GetKeysChanges(ctx, from, to)
Query users with recent device key updates.

Gets a list of users who have updated their device identity keys since a previous sync token.  The server should include in the results any users who:  * currently share a room with the calling user (ie, both users have   membership state ``join``); *and* * added new device identity keys or removed an existing device with   identity keys, between ``from`` and ``to``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **string**| The desired start point of the list. Should be the &#x60;&#x60;next_batch&#x60;&#x60; field from a response to an earlier call to |/sync|. Users who have not uploaded new device identity keys since this point, nor deleted existing devices with identity keys since then, will be excluded from the results. | 
**to** | **string**| The desired end point of the list. Should be the &#x60;&#x60;next_batch&#x60;&#x60; field from a recent call to |/sync| - typically the most recent such call. This may be used by the server as a hint to check its caches are up to date. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryKeys

> InlineResponse20017 QueryKeys(ctx, optional)
Download device identity keys.

Returns the current devices and identity keys for the given users.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**optional.Interface of InlineObject12**](InlineObject12.md)|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadKeys

> InlineResponse20018 UploadKeys(ctx, optional)
Upload end-to-end encryption keys.

Publishes end-to-end encryption keys for the device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keys** | [**optional.Interface of InlineObject13**](InlineObject13.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

