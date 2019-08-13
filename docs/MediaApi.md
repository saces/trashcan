# \MediaApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](MediaApi.md#GetConfig) | **Get** /_matrix/media/unstable/config | Get the configuration for the content repository.
[**GetContent**](MediaApi.md#GetContent) | **Get** /_matrix/media/unstable/download/{serverName}/{mediaId} | Download content from the content repository.
[**GetContentOverrideName**](MediaApi.md#GetContentOverrideName) | **Get** /_matrix/media/unstable/download/{serverName}/{mediaId}/{fileName} | Download content from the content repository. This is the same as the download endpoint above, except permitting a desired file name.
[**GetContentThumbnail**](MediaApi.md#GetContentThumbnail) | **Get** /_matrix/media/unstable/thumbnail/{serverName}/{mediaId} | Download a thumbnail of content from the content repository. See the &#x60;thumbnailing &lt;#thumbnails&gt;&#x60;_ section for more information.
[**GetUrlPreview**](MediaApi.md#GetUrlPreview) | **Get** /_matrix/media/unstable/preview_url | Get information about a URL for a client
[**UploadContent**](MediaApi.md#UploadContent) | **Post** /_matrix/media/unstable/upload | Upload some content to the content repository.



## GetConfig

> InlineResponse20047 GetConfig(ctx, )
Get the configuration for the content repository.

This endpoint allows clients to retrieve the configuration of the content repository, such as upload limitations. Clients SHOULD use this as a guide when using content repository endpoints. All values are intentionally left optional. Clients SHOULD follow the advice given in the field description when the field is not available.  **NOTE:** Both clients and server administrators should be aware that proxies between the client and the server may affect the apparent behaviour of content repository APIs, for example, proxies may enforce a lower upload size limit than is advertised by the server on this endpoint.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContent

> *os.File GetContent(ctx, serverName, mediaId, optional)
Download content from the content repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string**| The server name from the &#x60;&#x60;mxc://&#x60;&#x60; URI (the authoritory component)  | 
**mediaId** | **string**| The media ID from the &#x60;&#x60;mxc://&#x60;&#x60; URI (the path component)  | 
 **optional** | ***GetContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allowRemote** | **optional.Bool**| Indicates to the server that it should not attempt to fetch the media if it is deemed remote. This is to prevent routing loops where the server contacts itself. Defaults to true if not provided.  | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentOverrideName

> *os.File GetContentOverrideName(ctx, serverName, mediaId, fileName, optional)
Download content from the content repository. This is the same as the download endpoint above, except permitting a desired file name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string**| The server name from the &#x60;&#x60;mxc://&#x60;&#x60; URI (the authoritory component)  | 
**mediaId** | **string**| The media ID from the &#x60;&#x60;mxc://&#x60;&#x60; URI (the path component)  | 
**fileName** | **string**| A filename to give in the &#x60;&#x60;Content-Disposition&#x60;&#x60; header. | 
 **optional** | ***GetContentOverrideNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContentOverrideNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowRemote** | **optional.Bool**| Indicates to the server that it should not attempt to fetch the media if it is deemed remote. This is to prevent routing loops where the server contacts itself. Defaults to true if not provided.  | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentThumbnail

> *os.File GetContentThumbnail(ctx, serverName, mediaId, width, height, optional)
Download a thumbnail of content from the content repository. See the `thumbnailing <#thumbnails>`_ section for more information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverName** | **string**| The server name from the &#x60;&#x60;mxc://&#x60;&#x60; URI (the authoritory component)  | 
**mediaId** | **string**| The media ID from the &#x60;&#x60;mxc://&#x60;&#x60; URI (the path component)  | 
**width** | **int32**| The *desired* width of the thumbnail. The actual thumbnail may be larger than the size specified. | 
**height** | **int32**| The *desired* height of the thumbnail. The actual thumbnail may be larger than the size specified. | 
 **optional** | ***GetContentThumbnailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContentThumbnailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **method** | **optional.String**| The desired resizing method. See the &#x60;thumbnailing &lt;#thumbnails&gt;&#x60;_ section for more information. | 
 **allowRemote** | **optional.Bool**| Indicates to the server that it should not attempt to fetch the media if it is deemed remote. This is to prevent routing loops where the server contacts itself. Defaults to true if not provided.  | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, image/png, application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUrlPreview

> InlineResponse20048 GetUrlPreview(ctx, url, optional)
Get information about a URL for a client

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**| The URL to get a preview of. | 
 **optional** | ***GetUrlPreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUrlPreviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ts** | **optional.Int64**| The preferred point in time to return a preview for. The server may return a newer version if it does not have the requested version available. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadContent

> InlineResponse20049 UploadContent(ctx, content, optional)
Upload some content to the content repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**content** | **string**| The content to be uploaded. | 
 **optional** | ***UploadContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **optional.String**| The content type of the file being uploaded | 
 **filename** | **optional.String**| The name of the file being uploaded | 

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

