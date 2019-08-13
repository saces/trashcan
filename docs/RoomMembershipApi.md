# \RoomMembershipApi

All URIs are relative to *https://matrix.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Ban**](RoomMembershipApi.md#Ban) | **Post** /_matrix/client/unstable/rooms/{roomId}/ban | Ban a user in the room.
[**ForgetRoom**](RoomMembershipApi.md#ForgetRoom) | **Post** /_matrix/client/unstable/rooms/{roomId}/forget | Stop the requesting user remembering about a particular room.
[**GetJoinedRooms**](RoomMembershipApi.md#GetJoinedRooms) | **Get** /_matrix/client/unstable/joined_rooms | Lists the user&#39;s current rooms.
[**InviteBy3PID**](RoomMembershipApi.md#InviteBy3PID) | **Post** /_matrix/client/unstable/rooms/{roomId}/invite | Invite a user to participate in a particular room.
[**InviteUser**](RoomMembershipApi.md#InviteUser) | **Post** /_matrix/client/unstable/rooms/{roomId}/invite  | Invite a user to participate in a particular room.
[**JoinRoom**](RoomMembershipApi.md#JoinRoom) | **Post** /_matrix/client/unstable/join/{roomIdOrAlias} | Start the requesting user participating in a particular room.
[**JoinRoomById**](RoomMembershipApi.md#JoinRoomById) | **Post** /_matrix/client/unstable/rooms/{roomId}/join | Start the requesting user participating in a particular room.
[**Kick**](RoomMembershipApi.md#Kick) | **Post** /_matrix/client/unstable/rooms/{roomId}/kick | Kick a user from the room.
[**LeaveRoom**](RoomMembershipApi.md#LeaveRoom) | **Post** /_matrix/client/unstable/rooms/{roomId}/leave | Stop the requesting user participating in a particular room.
[**Unban**](RoomMembershipApi.md#Unban) | **Post** /_matrix/client/unstable/rooms/{roomId}/unban | Unban a user from the room.



## Ban

> map[string]interface{} Ban(ctx, roomId, body)
Ban a user in the room.

Ban a user in the room. If the user is currently in the room, also kick them.  When a user is banned from a room, they may not join it or be invited to it until they are unbanned.  The caller must have the required power level in order to perform this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier (not alias) from which the user should be banned. | 
**body** | [**InlineObject24**](InlineObject24.md)|  | 

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


## ForgetRoom

> map[string]interface{} ForgetRoom(ctx, roomId)
Stop the requesting user remembering about a particular room.

This API stops a user remembering about a particular room.  In general, history is a first class citizen in Matrix. After this API is called, however, a user will no longer be able to retrieve history for this room. If all users on a homeserver forget a room, the room is eligible for deletion from that homeserver.  If the user is currently joined to the room, they must leave the room before calling this API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier to forget. | 

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


## GetJoinedRooms

> InlineResponse20014 GetJoinedRooms(ctx, )
Lists the user's current rooms.

This API returns a list of the user's current rooms.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteBy3PID

> map[string]interface{} InviteBy3PID(ctx, roomId, body)
Invite a user to participate in a particular room.

.. _invite-by-third-party-id-endpoint:  *Note that there are two forms of this API, which are documented separately. This version of the API does not require that the inviter know the Matrix identifier of the invitee, and instead relies on third party identifiers. The homeserver uses an identity server to perform the mapping from third party identifier to a Matrix identifier. The other is documented in the* `joining rooms section`_.  This API invites a user to participate in a particular room. They do not start participating in the room until they actually join the room.  Only users currently in a particular room can invite other users to join that room.  If the identity server did know the Matrix user identifier for the third party identifier, the homeserver will append a ``m.room.member`` event to the room.  If the identity server does not know a Matrix user identifier for the passed third party identifier, the homeserver will issue an invitation which can be accepted upon providing proof of ownership of the third party identifier. This is achieved by the identity server generating a token, which it gives to the inviting homeserver. The homeserver will add an ``m.room.third_party_invite`` event into the graph for the room, containing that token.  When the invitee binds the invited third party identifier to a Matrix user ID, the identity server will give the user a list of pending invitations, each containing:  - The room ID to which they were invited  - The token given to the homeserver  - A signature of the token, signed with the identity server's private key  - The matrix user ID who invited them to the room  If a token is requested from the identity server, the homeserver will append a ``m.room.third_party_invite`` event to the room.  .. _joining rooms section: `invite-by-user-id-endpoint`_

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier (not alias) to which to invite the user. | 
**body** | [**InlineObject25**](InlineObject25.md)|  | 

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


## InviteUser

> map[string]interface{} InviteUser(ctx, roomId, body)
Invite a user to participate in a particular room.

.. _invite-by-user-id-endpoint:  *Note that there are two forms of this API, which are documented separately. This version of the API requires that the inviter knows the Matrix identifier of the invitee. The other is documented in the* `third party invites section`_.  This API invites a user to participate in a particular room. They do not start participating in the room until they actually join the room.  Only users currently in a particular room can invite other users to join that room.  If the user was invited to the room, the homeserver will append a ``m.room.member`` event to the room.  .. _third party invites section: `invite-by-third-party-id-endpoint`_

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier (not alias) to which to invite the user. | 
**body** | [**InlineObject26**](InlineObject26.md)|  | 

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


## JoinRoom

> InlineResponse20013 JoinRoom(ctx, roomIdOrAlias, optional)
Start the requesting user participating in a particular room.

*Note that this API takes either a room ID or alias, unlike* ``/room/{roomId}/join``.  This API starts a user participating in a particular room, if that user is allowed to participate in that room. After this call, the client is allowed to see all current state events in the room, and all subsequent events associated with the room until the user leaves the room.  After a user has joined a room, the room will appear as an entry in the response of the |/initialSync|_ and |/sync|_ APIs.  If a ``third_party_signed`` was supplied, the homeserver must verify that it matches a pending ``m.room.third_party_invite`` event in the room, and perform key validity checking if required by the event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomIdOrAlias** | **string**| The room identifier or alias to join. | 
 **optional** | ***JoinRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JoinRoomOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverName** | [**optional.Interface of []string**](string.md)| The servers to attempt to join the room through. One of the servers must be participating in the room. | 
 **thirdPartySigned** | [**optional.Interface of InlineObject10**](InlineObject10.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinRoomById

> InlineResponse20013 JoinRoomById(ctx, roomId, optional)
Start the requesting user participating in a particular room.

*Note that this API requires a room ID, not alias.* ``/join/{roomIdOrAlias}`` *exists if you have a room alias.*  This API starts a user participating in a particular room, if that user is allowed to participate in that room. After this call, the client is allowed to see all current state events in the room, and all subsequent events associated with the room until the user leaves the room.  After a user has joined a room, the room will appear as an entry in the response of the |/initialSync|_ and |/sync|_ APIs.  If a ``third_party_signed`` was supplied, the homeserver must verify that it matches a pending ``m.room.third_party_invite`` event in the room, and perform key validity checking if required by the event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier (not alias) to join. | 
 **optional** | ***JoinRoomByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JoinRoomByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thirdPartySigned** | [**optional.Interface of InlineObject27**](InlineObject27.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Kick

> map[string]interface{} Kick(ctx, roomId, body)
Kick a user from the room.

Kick a user from the room.  The caller must have the required power level in order to perform this operation.  Kicking a user adjusts the target member's membership state to be ``leave`` with an optional ``reason``. Like with other membership changes, a user can directly adjust the target member's state by making a request to ``/rooms/<room id>/state/m.room.member/<user id>``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier (not alias) from which the user should be kicked. | 
**body** | [**InlineObject28**](InlineObject28.md)|  | 

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


## LeaveRoom

> map[string]interface{} LeaveRoom(ctx, roomId)
Stop the requesting user participating in a particular room.

This API stops a user participating in a particular room.  If the user was already in the room, they will no longer be able to see new events in the room. If the room requires an invite to join, they will need to be re-invited before they can re-join.  If the user was invited to the room, but had not joined, this call serves to reject the invite.  The user will still be allowed to retrieve history from the room which they were previously allowed to see.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier to leave. | 

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


## Unban

> map[string]interface{} Unban(ctx, roomId, body)
Unban a user from the room.

Unban a user from the room. This allows them to be invited to the room, and join if they would otherwise be allowed to join according to its join rules.  The caller must have the required power level in order to perform this operation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **string**| The room identifier (not alias) from which the user should be unbanned. | 
**body** | [**InlineObject33**](InlineObject33.md)|  | 

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

