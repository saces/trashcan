# \DeviceManagementApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevice**](DeviceManagementApi.md#DeleteDevice) | **Delete** /_matrix/client/unstable/devices/{deviceId} | Delete a device
[**DeleteDevices**](DeviceManagementApi.md#DeleteDevices) | **Post** /_matrix/client/unstable/delete_devices | Bulk deletion of devices
[**GetDevice**](DeviceManagementApi.md#GetDevice) | **Get** /_matrix/client/unstable/devices/{deviceId} | Get a single device
[**GetDevices**](DeviceManagementApi.md#GetDevices) | **Get** /_matrix/client/unstable/devices | List registered devices for the current user
[**UpdateDevice**](DeviceManagementApi.md#UpdateDevice) | **Put** /_matrix/client/unstable/devices/{deviceId} | Update a device



## DeleteDevice

> map[string]interface{} DeleteDevice(ctx, deviceId, optional)
Delete a device

This API endpoint uses the `User-Interactive Authentication API`_.  Deletes the given device, and invalidates any access token associated with it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| The device to delete. | 
 **optional** | ***DeleteDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InlineObject7**](InlineObject7.md)|  | 

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


## DeleteDevices

> map[string]interface{} DeleteDevices(ctx, optional)
Bulk deletion of devices

This API endpoint uses the `User-Interactive Authentication API`_.  Deletes the given devices, and invalidates any access token associated with them.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

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


## GetDevice

> InlineResponse2008 GetDevice(ctx, deviceId)
Get a single device

Gets information on a single device, by device id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| The device to retrieve. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> InlineResponse2007 GetDevices(ctx, )
List registered devices for the current user

Gets information about all devices for the current user.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> map[string]interface{} UpdateDevice(ctx, deviceId, body)
Update a device

Updates the metadata on the given device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| The device to update. | 
**body** | [**InlineObject6**](InlineObject6.md)|  | 

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

