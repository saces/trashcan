# \UserDataApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](UserDataApi.md#ChangePassword) | **Post** /_matrix/client/unstable/account/password | Changes a user&#39;s password.
[**CheckUsernameAvailability**](UserDataApi.md#CheckUsernameAvailability) | **Get** /_matrix/client/unstable/register/available | Checks to see if a username is available on the server.
[**DeactivateAccount**](UserDataApi.md#DeactivateAccount) | **Post** /_matrix/client/unstable/account/deactivate | Deactivate a user&#39;s account.
[**Delete3pidFromAccount**](UserDataApi.md#Delete3pidFromAccount) | **Post** /_matrix/client/unstable/account/3pid/delete | Deletes a third party identifier from the user&#39;s account
[**DeleteRoomTag**](UserDataApi.md#DeleteRoomTag) | **Delete** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags/{tag} | Remove a tag from the room.
[**GetAccount3PIDs**](UserDataApi.md#GetAccount3PIDs) | **Get** /_matrix/client/unstable/account/3pid | Gets a list of a user&#39;s third party identifiers.
[**GetAccountData**](UserDataApi.md#GetAccountData) | **Get** /_matrix/client/unstable/user/{userId}/account_data/{type} | Get some account_data for the user.
[**GetAccountDataPerRoom**](UserDataApi.md#GetAccountDataPerRoom) | **Get** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/account_data/{type} | Get some account_data for the user.
[**GetAvatarUrl**](UserDataApi.md#GetAvatarUrl) | **Get** /_matrix/client/unstable/profile/{userId}/avatar_url | Get the user&#39;s avatar URL.
[**GetDisplayName**](UserDataApi.md#GetDisplayName) | **Get** /_matrix/client/unstable/profile/{userId}/displayname | Get the user&#39;s display name.
[**GetRoomTags**](UserDataApi.md#GetRoomTags) | **Get** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags | List the tags for a room.
[**GetTokenOwner**](UserDataApi.md#GetTokenOwner) | **Get** /_matrix/client/unstable/account/whoami | Gets information about the owner of an access token.
[**GetUserProfile**](UserDataApi.md#GetUserProfile) | **Get** /_matrix/client/unstable/profile/{userId} | Get this user&#39;s profile information.
[**Post3PIDs**](UserDataApi.md#Post3PIDs) | **Post** /_matrix/client/unstable/account/3pid | Adds contact information to the user&#39;s account.
[**Register**](UserDataApi.md#Register) | **Post** /_matrix/client/unstable/register | Register for an account on this homeserver.
[**SearchUserDirectory**](UserDataApi.md#SearchUserDirectory) | **Post** /_matrix/client/unstable/user_directory/search | Searches the user directory.
[**SetAccountData**](UserDataApi.md#SetAccountData) | **Put** /_matrix/client/unstable/user/{userId}/account_data/{type} | Set some account_data for the user.
[**SetAccountDataPerRoom**](UserDataApi.md#SetAccountDataPerRoom) | **Put** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/account_data/{type} | Set some account_data for the user.
[**SetAvatarUrl**](UserDataApi.md#SetAvatarUrl) | **Put** /_matrix/client/unstable/profile/{userId}/avatar_url | Set the user&#39;s avatar URL.
[**SetDisplayName**](UserDataApi.md#SetDisplayName) | **Put** /_matrix/client/unstable/profile/{userId}/displayname | Set the user&#39;s display name.
[**SetRoomTag**](UserDataApi.md#SetRoomTag) | **Put** /_matrix/client/unstable/user/{userId}/rooms/{roomId}/tags/{tag} | Add a tag to a room.



## ChangePassword

> map[string]interface{} ChangePassword(ctx, optional)
Changes a user's password.

Changes the password for an account on this homeserver.  This API endpoint uses the `User-Interactive Authentication API`_ to ensure the user changing the password is actually the owner of the account.  An access token should be submitted to this endpoint if the client has an active session.  The homeserver may change the flows available depending on whether a valid access token is provided. The homeserver SHOULD NOT revoke the access token provided in the request, however all other access tokens for the user should be revoked if the request succeeds.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChangePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChangePasswordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject3**](InlineObject3.md)|  | 

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


## CheckUsernameAvailability

> InlineResponse20032 CheckUsernameAvailability(ctx, username)
Checks to see if a username is available on the server.

Checks to see if a username is available, and valid, for the server.  The server should check to ensure that, at the time of the request, the username requested is available for use. This includes verifying that an application service has not claimed the username and that the username fits the server's desired requirements (for example, a server could dictate that it does not permit usernames with underscores).  Matrix clients may wish to use this API prior to attempting registration, however the clients must also be aware that using this API does not normally reserve the username. This can mean that the username becomes unavailable between checking its availability and attempting to register it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**| The username to check the availability of. | [default to my_cool_localpart]

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAccount

> InlineResponse2002 DeactivateAccount(ctx, optional)
Deactivate a user's account.

Deactivate the user's account, removing all ability for the user to login again.  This API endpoint uses the `User-Interactive Authentication API`_.  An access token should be submitted to this endpoint if the client has an active session.  The homeserver may change the flows available depending on whether a valid access token is provided.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeactivateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeactivateAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete3pidFromAccount

> InlineResponse2001 Delete3pidFromAccount(ctx, optional)
Deletes a third party identifier from the user's account

Removes a third party identifier from the user's account. This might not cause an unbind of the identifier from the identity server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Delete3pidFromAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Delete3pidFromAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomTag

> map[string]interface{} DeleteRoomTag(ctx, userId, roomId, tag)
Remove a tag from the room.

Remove a tag from the room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The id of the user to remove a tag for. The access token must be authorized to make requests for this user ID. | 
**roomId** | **string**| The ID of the room to remove a tag from. | 
**tag** | **string**| The tag to remove. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount3PIDs

> InlineResponse200 GetAccount3PIDs(ctx, )
Gets a list of a user's third party identifiers.

Gets a list of the third party identifiers that the homeserver has associated with the user's account.  This is *not* the same as the list of third party identifiers bound to the user's Matrix ID in identity servers.  Identifiers in this list may be used by the homeserver as, for example, identifiers that it will accept to reset the user's account password.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountData

> map[string]interface{} GetAccountData(ctx, userId, type_)
Get some account_data for the user.

Get some account_data for the client. This config is only visible to the user that set the account_data.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user to get account_data for. The access token must be authorized to make requests for this user ID. | 
**type_** | **string**| The event type of the account_data to get. Custom types should be namespaced to avoid clashes. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountDataPerRoom

> map[string]interface{} GetAccountDataPerRoom(ctx, userId, roomId, type_)
Get some account_data for the user.

Get some account_data for the client on a given room. This config is only visible to the user that set the account_data.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user to set account_data for. The access token must be authorized to make requests for this user ID. | 
**roomId** | **string**| The ID of the room to get account_data for. | 
**type_** | **string**| The event type of the account_data to get. Custom types should be namespaced to avoid clashes. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatarUrl

> InlineResponse20024 GetAvatarUrl(ctx, userId)
Get the user's avatar URL.

Get the user's avatar URL. This API may be used to fetch the user's own avatar URL or to query the URL of other users; either locally or on remote homeservers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose avatar URL to get. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisplayName

> InlineResponse20025 GetDisplayName(ctx, userId)
Get the user's display name.

Get the user's display name. This API may be used to fetch the user's own displayname or to query the name of other users; either locally or on remote homeservers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose display name to get. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomTags

> InlineResponse20043 GetRoomTags(ctx, userId, roomId)
List the tags for a room.

List the tags set by a user on a room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The id of the user to get tags for. The access token must be authorized to make requests for this user ID. | 
**roomId** | **string**| The ID of the room to get tags for. | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenOwner

> InlineResponse2003 GetTokenOwner(ctx, )
Gets information about the owner of an access token.

Gets information about the owner of a given access token.   Note that, as with the rest of the Client-Server API,  Application Services may masquerade as users within their  namespace by giving a ``user_id`` query parameter. In this  situation, the server should verify that the given ``user_id`` is registered by the appservice, and return it in the response  body.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserProfile

> InlineResponse20023 GetUserProfile(ctx, userId)
Get this user's profile information.

Get the combined profile information for this user. This API may be used to fetch the user's own profile information or other users; either locally or on remote homeservers. This API may return keys which are not limited to ``displayname`` or ``avatar_url``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose profile information to get. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post3PIDs

> Post3PIDs(ctx, optional)
Adds contact information to the user's account.

Adds contact information to the user's account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Post3PIDsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Post3PIDsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, schema

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> InlineResponse20031 Register(ctx, optional)
Register for an account on this homeserver.

This API endpoint uses the `User-Interactive Authentication API`_, except in the cases where a guest account is being registered.  Register for an account on this homeserver.  There are two kinds of user account:  - `user` accounts. These accounts may use the full API described in this specification.  - `guest` accounts. These accounts may have limited permissions and may not be supported by all servers.  If registration is successful, this endpoint will issue an access token the client can use to authorize itself in subsequent requests.  If the client does not supply a ``device_id``, the server must auto-generate one.  The server SHOULD register an account with a User ID based on the ``username`` provided, if any. Note that the grammar of Matrix User ID localparts is restricted, so the server MUST either map the provided ``username`` onto a ``user_id`` in a logical manner, or reject ``username``\\s which do not comply to the grammar, with ``M_INVALID_USERNAME``.  Matrix clients MUST NOT assume that localpart of the registered ``user_id`` matches the provided ``username``.  The returned access token must be associated with the ``device_id`` supplied by the client or generated by the server. The server may invalidate any access token previously associated with that device. See `Relationship between access tokens and devices`_.  When registering a guest account, all parameters in the request body with the exception of ``initial_device_display_name`` MUST BE ignored by the server. The server MUST pick a ``device_id`` for the account regardless of input.  Any user ID returned by this API must conform to the grammar given in the `Matrix specification <../appendices.html#user-identifiers>`_.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **optional.String**| The kind of account to register. Defaults to &#x60;&#x60;user&#x60;&#x60;. | [default to user]
 **body** | [**optional.Interface of InlineObject23**](InlineObject23.md)|  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUserDirectory

> InlineResponse20044 SearchUserDirectory(ctx, optional)
Searches the user directory.

Performs a search for users on the homeserver. The homeserver may determine which subset of users are searched, however the homeserver MUST at a minimum consider the users the requesting user shares a room with and those who reside in public rooms (known to the homeserver). The search MUST consider local users to the homeserver, and SHOULD query remote users as part of the search.  The search is performed case-insensitively on user IDs and display names preferably using a collation determined based upon the  ``Accept-Language`` header provided in the request, if present.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchUserDirectoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUserDirectoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InlineObject37**](InlineObject37.md)|  | 

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccountData

> SetAccountData(ctx, userId, type_, content)
Set some account_data for the user.

Set some account_data for the client. This config is only visible to the user that set the account_data. The config will be synced to clients in the top-level ``account_data``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user to set account_data for. The access token must be authorized to make requests for this user ID. | 
**type_** | **string**| The event type of the account_data to set. Custom types should be namespaced to avoid clashes. | 
**content** | **map[string]interface{}**| The content of the account_data | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccountDataPerRoom

> SetAccountDataPerRoom(ctx, userId, roomId, type_, content)
Set some account_data for the user.

Set some account_data for the client on a given room. This config is only visible to the user that set the account_data. The config will be synced to clients in the per-room ``account_data``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The ID of the user to set account_data for. The access token must be authorized to make requests for this user ID. | 
**roomId** | **string**| The ID of the room to set account_data on. | 
**type_** | **string**| The event type of the account_data to set. Custom types should be namespaced to avoid clashes. | 
**content** | **map[string]interface{}**| The content of the account_data | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAvatarUrl

> map[string]interface{} SetAvatarUrl(ctx, userId, avatarUrl)
Set the user's avatar URL.

This API sets the given user's avatar URL. You must have permission to set this user's avatar URL, e.g. you need to have their ``access_token``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose avatar URL to set. | 
**avatarUrl** | [**InlineObject16**](InlineObject16.md)|  | 

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


## SetDisplayName

> map[string]interface{} SetDisplayName(ctx, userId, displayName)
Set the user's display name.

This API sets the given user's display name. You must have permission to set this user's display name, e.g. you need to have their ``access_token``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user whose display name to set. | 
**displayName** | [**InlineObject17**](InlineObject17.md)|  | 

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


## SetRoomTag

> map[string]interface{} SetRoomTag(ctx, userId, roomId, tag, body)
Add a tag to a room.

Add a tag to the room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The id of the user to add a tag for. The access token must be authorized to make requests for this user ID. | 
**roomId** | **string**| The ID of the room to add a tag to. | 
**tag** | **string**| The tag to add. | 
**body** | [**InlineObject36**](InlineObject36.md)|  | 

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

