# \OpenIDApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestOpenIdToken**](OpenIDApi.md#RequestOpenIdToken) | **Post** /_matrix/client/unstable/user/{userId}/openid/request_token | Get an OpenID token object to verify the requester&#39;s identity.



## RequestOpenIdToken

> InlineResponse20042 RequestOpenIdToken(ctx, userId, body)
Get an OpenID token object to verify the requester's identity.

Gets an OpenID token object that the requester may supply to another service to verify their identity in Matrix. The generated token is only valid for exchanging for user information from the federation API for OpenID.  The access token generated is only valid for the OpenID API. It cannot be used to request another OpenID access token or call ``/sync``, for example.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user to request and OpenID token for. Should be the user who is authenticated for the request. | 
**body** | **map[string]interface{}**| An empty object. Reserved for future expansion. | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

