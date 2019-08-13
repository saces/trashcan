# \ServerAdministrationApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVersions**](ServerAdministrationApi.md#GetVersions) | **Get** /_matrix/client/versions | Gets the versions of the specification supported by the server.
[**GetWellknown**](ServerAdministrationApi.md#GetWellknown) | **Get** /.well-known/matrix/client | Gets Matrix server discovery information about the domain.
[**GetWhoIs**](ServerAdministrationApi.md#GetWhoIs) | **Get** /_matrix/client/unstable/admin/whois/{userId} | Gets information about a particular user.



## GetVersions

> InlineResponse20046 GetVersions(ctx, )
Gets the versions of the specification supported by the server.

Gets the versions of the specification supported by the server.  Values will take the form ``rX.Y.Z``.  Only the latest ``Z`` value will be reported for each supported ``X.Y`` value. i.e. if the server implements ``r0.0.0``, ``r0.0.1``, and ``r1.2.0``, it will report ``r0.0.1`` and ``r1.2.0``.  The server may additionally advertise experimental features it supports through ``unstable_features``. These features should be namespaced and may optionally include version information within their name if desired. Features listed here are not for optionally toggling parts of the Matrix specification and should only be used to advertise support for a feature which has not yet landed in the spec. For example, a feature currently undergoing the proposal process may appear here and eventually be taken off this list once the feature lands in the spec and the server deems it reasonable to do so. Servers may wish to keep advertising features here after they've been released into the spec to give clients a chance to upgrade appropriately. Additionally, clients should avoid using unstable features in their stable releases.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellknown

> map[string]map[string]interface{} GetWellknown(ctx, )
Gets Matrix server discovery information about the domain.

Gets discovery information about the domain. The file may include additional keys, which MUST follow the Java package naming convention, e.g. ``com.example.myapp.property``. This ensures property names are suitably namespaced for each application and reduces the risk of clashes.  Note that this endpoint is not necessarily handled by the homeserver, but by another webserver, to be used for discovering the homeserver URL.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhoIs

> InlineResponse2004 GetWhoIs(ctx, userId)
Gets information about a particular user.

Gets information about a particular user.  This API may be restricted to only be called by the user being looked up, or by a server admin. Server-local administrator privileges are not specified in this document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user to look up. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

